package repository

import (
	"context"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

type githubClient struct {
	client *github.Client
}

// NewGithubRESTClient creates a GitHub REST API client wrapper
func NewGithubRESTClient(config *entity.Configuration) internal.GithubClient {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)

	oauth := oauth2.NewClient(ctx, ts)
	github := github.NewClient(oauth)

	github.BaseURL = config.BaseURL

	return &githubClient{client: github}
}

func (cl *githubClient) GetPullRequest(
	ctx context.Context,
	owner string,
	name string,
	event int,
) (*entity.PullRequest, error) {
	pullRequest, _, err := cl.client.PullRequests.Get(ctx, owner, name, event)

	return &entity.PullRequest{
		Title: pullRequest.GetTitle(),
		Body: pullRequest.GetBody(),
		IsDraft: pullRequest.GetDraft(),
		Branch: pullRequest.GetHead().GetRef(),
		Changes: pullRequest.GetChangedFiles(),
		Author: entity.Actor{
			Name: pullRequest.GetUser().GetName(),
			Type: pullRequest.GetUser().GetType(),
		},
	}, err
}

func (cl *githubClient) GetUser(ctx context.Context, name string) (*github.User, error) {
	user, _, err := cl.client.Users.Get(ctx, name)

	return user, err
}

func (cl *githubClient) GetIssue(
	ctx context.Context,
	owner string,
	name string,
	number int,
) (*github.Issue, error) {
	issue, _, err := cl.client.Issues.Get(ctx, owner, name, number)

	return issue, err
}

func (cl *githubClient) GetPermissionLevel(
	ctx context.Context,
	owner string,
	name string,
	user string,
) (*github.RepositoryPermissionLevel, error) {
	perms, _, err := cl.client.Repositories.GetPermissionLevel(ctx, owner, name, user)

	return perms, err
}

func (cl *githubClient) GetComments(
	ctx context.Context,
	owner string,
	name string,
	number int,
) ([]*github.IssueComment, error) {
	comments, _, err := cl.client.Issues.ListComments(
		ctx,
		owner,
		name,
		number,
		&github.IssueListCommentsOptions{},
	)

	return comments, err
}

func (cl *githubClient) GetCommits(
	ctx context.Context,
	owner string,
	name string,
	event int,
) ([]*github.RepositoryCommit, error) {
	commits, _, err := cl.client.PullRequests.ListCommits(ctx, owner, name, event, nil)

	return commits, err
}

func (cl *githubClient) EditComment(
	ctx context.Context,
	owner string,
	name string,
	id int64,
	comment *github.IssueComment,
) error {
	_, _, err := cl.client.Issues.EditComment(
		ctx,
		owner,
		name,
		id,
		comment,
	)

	return err
}

func (cl *githubClient) CreateComment(
	ctx context.Context,
	owner string,
	name string,
	event int,
	comment *github.IssueComment,
) error {
	_, _, err := cl.client.Issues.CreateComment(ctx, owner, name, event, comment)

	return err
}

func (cl *githubClient) Label(
	ctx context.Context,
	owner string,
	name string,
	event int,
	label string,
) error {
	_, _, err := cl.client.Issues.AddLabelsToIssue(ctx, owner, name, event, []string{
		label,
	})

	return err
}

func (cl *githubClient) Close(
	ctx context.Context,
	owner string,
	name string,
	event int,
) error {
	state := constants.Closed

	_, _, err := cl.client.PullRequests.Edit(ctx, owner, name, event, &github.PullRequest{
		State: &state,
	})

	return err
}