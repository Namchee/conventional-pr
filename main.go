package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/utils"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func main() {
	config, err := entity.ReadConfig()

	if err != nil {
		log.Fatalln(err)
	}

	meta, _ := entity.CreateMeta(
		utils.ReadEnvString("GITHUB_REPOSITORY"),
	)
	event, _ := entity.ReadEvent(
		utils.ReadEnvString("GITHUB_EVENT_PATH"),
	)

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	pullRequest, _, err := client.PullRequests.Get(
		ctx,
		meta.Owner,
		meta.Name,
		event.Number,
	)
	commitList, _, err := client.PullRequests.ListCommits(
		ctx,
		meta.Owner,
		meta.Name,
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

	if !contains(constants.Events, event.Action) { // ignore most events
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

	log.Printf("Title is valid: %s", strconv.FormatBool(isTitleValid))
	log.Printf("Length of body: %d", len(body))
	log.Printf("Count of issues: %d", len(issueMentions))
	log.Printf("Too many changes in PR: %s", strconv.FormatBool(hasTooManyChanges))
	if badCommit != nil {
		log.Printf("No bad commits found.")
	} else {
		log.Printf("Bad commits found.")
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
