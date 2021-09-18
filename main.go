package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

const (
	closedState = "closed"
)

// Metadata for the current repository
type Metadata struct {
	name  string
	owner string
}

// Options passed from the action input
type Options struct {
	token         string
	checkDraft    bool
	label         string
	strict        bool
	close         bool
	assignee      bool
	template      string
	allowedTypes  []string
	allowedScopes []string
	maxFileChange int
	issue         bool
	ignoreBot     bool
}

// Event that triggers the action
type Event struct {
	Action string "json:action"
	Number int    "json:number"
}

// BadCommit that doesn't follow the conventional commit message
type BadCommit struct {
	url     string
	message string
}

// Get all the required environment variables and encapsulate them
// in a custom Options struct.
func getOptionsValues() Options {
	token := os.Getenv("INPUT_ACCESS_TOKEN")

	if len(token) == 0 {
		log.Fatalln("Missing GitHub Access Token")
	}

	label := os.Getenv("INPUT_LABEL")
	close := os.Getenv("INPUT_CLOSE") == "true"
	assignee := os.Getenv("INPUT_ASSIGNEE") == "true"
	template := os.Getenv("INPUT_TEMPLATE")
	strict := os.Getenv("INPUT_STRICT") == "true"
	allowedTypes := strings.Split(os.Getenv("INPUT_ALLOWED_TYPES"), ",")
	allowedScopes := strings.Split(os.Getenv("INPUT_ALLOWED_SCOPES"), ",")
	checkDraft := os.Getenv("INPUT_CHECK_DRAFT") == "true"
	maxFileChange, _ := strconv.Atoi(os.Getenv("INPUT_MAXIMUM_FILE_CHANGE"))
	issue := os.Getenv("INPUT_LINK_ISSUE") == "true"
	ignoreBot := os.Getenv("INPUT_IGNORE_BOT") == "true"

	return Options{
		token,
		checkDraft,
		label,
		strict,
		close,
		assignee,
		template,
		allowedTypes,
		allowedScopes,
		maxFileChange,
		issue,
		ignoreBot,
	}
}

func main() {
	options := getOptionsValues()

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: options.token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	metadata := getRepositoryMetadata(os.Getenv("GITHUB_REPOSITORY"))

	if metadata == nil {
		log.Fatalln("Failed to read repository metadata")
	}

	file, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))

	if err != nil {
		log.Fatalln("Failed to open event data")
	}

	var event Event

	if err = json.NewDecoder(file).Decode(&event); err != nil {
		log.Fatalln("Failed to parse event data")
	}
	file.Close()

	pullRequest, _, err := client.PullRequests.Get(
		ctx,
		metadata.owner,
		metadata.name,
		event.Number,
	)
	commitList, _, err := client.PullRequests.ListCommits(
		ctx,
		metadata.owner,
		metadata.name,
		event.Number,
		nil,
	)
	body := pullRequest.GetBody()

	if err != nil {
		log.Fatalln("Failed to fetch pull request data")
	}

	user, _, err := client.Users.Get(ctx, pullRequest.GetUser().GetLogin())

	if err != nil {
		log.Fatalln("Failed to fetch user data")
	}

	events := []string{"opened", "reopened", "ready_for_review"} // probably will change

	if !contains(events, event.Action) { // ignore most events
		os.Exit(0)
	}

	privilege, _, err := client.Repositories.GetPermissionLevel(
		ctx,
		metadata.owner,
		metadata.name,
		user.GetLogin(),
	)

	if err != nil {
		log.Fatalln("Failed to fetch privilege level")
	}

	bypassByPrivilege := strings.ToLower(privilege.GetPermission()) == "admin" &&
		!options.strict

	if bypassByPrivilege {
		log.Println("Pull request checks are skipped due to high administrative privileges")
	}

	fileChangeCount := pullRequest.GetChangedFiles()
	isDraft := !options.checkDraft && pullRequest.GetDraft()

	if isDraft {
		log.Println("Pull request checks are skipped since the corresponding pull request is still a draft")
	}

	ignoreBot := strings.ToLower(user.GetType()) == "bot" &&
		options.ignoreBot

	if ignoreBot {
		log.Println("Pull request checks are skipped since the corresponding pull request is submitted by a bot")
	}

	if bypassByPrivilege || isDraft || ignoreBot {
		os.Exit(0)
	}

	var reason string

	isTitleValid := isConventional(pullRequest.GetTitle(), &options)
	issueMentions := getIssues(&ctx, client, metadata.owner, metadata.name, body)
	badCommit := getUnconventionalCommit(commitList, &options)
	hasTooManyChanges := options.maxFileChange > 0 &&
		fileChangeCount > options.maxFileChange

	if !isTitleValid { // Nonconventional
		reason = "Pull requests title must follow the [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/) style."
	} else if len(body) == 0 { // empty PR body
		reason = "Pull requests must have a clear and concise description."
	} else if len(issueMentions) == 0 && options.issue { // doesn't reference an issue
		reason = "Pull requests must at least refer to an issue."
	} else if hasTooManyChanges { // too many changed files
		reason = "This pull request has too many file changes. Consider splitting this pull request into some few smaller PRs."
	} else if badCommit != nil {
		reason = fmt.Sprintf("Commit message `%s` at commit %s doesn't follow the [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/) style.", badCommit.message, badCommit.url)
	}

	if len(reason) > 0 {
		err := closePullRequest(
			reason,
			&options,
			client,
			&ctx,
			metadata.owner,
			metadata.name,
			pullRequest.GetNumber(),
		)

		if err != nil {
			log.Printf(reason)
			log.Fatalf("Failed to change pull request: %s", err)
		}

		log.Fatalln(reason)
	} else {
		log.Printf("Pull request %d is a conventional PR!", pullRequest.GetNumber())
	}
}

// Utility function to check if `v` is present on array `s`
func contains(s []string, v string) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}

	return false
}

// Get repository owner and name from raw GitHub metadata.
func getRepositoryMetadata(rawMetadata string) *Metadata {
	tokens := strings.SplitN(rawMetadata, "/", 2)

	if len(tokens) != 2 {
		return nil
	}

	return &Metadata{tokens[1], tokens[0]}
}

// Determine whether a text follows conventional commit style.
func isConventional(
	text string,
	options *Options,
) bool {
	pattern := regexp.MustCompile(`([\w\s]+)(\(([\w\s]+)\))?!?: [\w\s]+`)

	if !pattern.Match([]byte(text)) {
		return false
	}

	submatches := pattern.FindStringSubmatch(text)

	if (len(options.allowedTypes) > 1 && !contains(options.allowedTypes, submatches[1])) ||
		len(options.allowedScopes) > 1 && !contains(options.allowedScopes, submatches[3]) {
		return false
	}

	return true
}

// Get all issue mentions from the pull request body.
func getIssueMentions(
	prBody string,
) []int {
	pattern := regexp.MustCompile("#(\\d+)")

	numberStrings := pattern.FindAllStringSubmatch(prBody, -1)
	var numbers []int

	if numberStrings != nil {
		for _, str := range numberStrings {
			intRep, _ := strconv.Atoi(str[1])
			numbers = append(numbers, intRep)
		}
	}

	return numbers
}

// Get the first unconventional commit message.
//
// Will return nil if not found.
func getUnconventionalCommit(
	commits []*github.RepositoryCommit,
	options *Options,
) *BadCommit {
	for _, repoCommit := range commits {
		commit := repoCommit.GetCommit()
		msg := commit.GetMessage()

		if !isConventional(msg, options) {
			return &BadCommit{
				commit.GetSHA(),
				msg,
			}
		}
	}

	return nil
}

// Get all issues from a repository.
func getIssues(
	ctx *context.Context,
	client *github.Client,
	owner string,
	repo string,
	prBody string,
) []*github.Issue {
	var issues []*github.Issue
	numbers := getIssueMentions(prBody)

	for _, number := range numbers {
		issue, _, err := client.Issues.Get(*ctx, owner, repo, number)

		if err == nil {
			issues = append(issues, issue)
		}
	}

	return issues
}

// Close a pull request with specific reason and options.
func closePullRequest(
	reason string,
	options *Options,
	client *github.Client,
	ctx *context.Context,
	owner string,
	repo string,
	number int,
) error {
	_, _, err := client.Issues.CreateComment(*ctx, owner, repo, number, &github.IssueComment{
		Body: github.String(fmt.Sprintf("%s\n\n**Reason**: %s", options.template, reason)),
	})
	if err != nil {
		return err
	}

	_, _, err = client.Issues.AddLabelsToIssue(*ctx, owner, repo, number, []string{
		options.label,
	})

	if err != nil {
		return err
	}

	if options.close {
		_, _, err := client.PullRequests.Edit(*ctx, owner, repo, number, &github.PullRequest{
			State: github.String(closedState),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
