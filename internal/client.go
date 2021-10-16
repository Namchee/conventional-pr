package internal

import (
	"context"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/google/go-github/v32/github"
)

// GithubClient handles all interaction with Github's API
// Designed this way for easier software testing
type GithubClient interface {
	GetUser(context.Context, string) (*github.User, error)
	GetIssue(context.Context, string, string, int) (*github.Issue, error)
	GetPermissionLevel(context.Context, string, string, string) (*github.RepositoryPermissionLevel, error)
	GetCommits(context.Context, string, string, int) ([]*github.RepositoryCommit, error)
	Comment(context.Context, string, string, int, *github.PullRequestComment) error
	Label(context.Context, string, string, int, string) error
	Close(context.Context, string, string, int) error
}

type githubClient struct {
	client *github.Client
}

func NewGithubClient(client *github.Client) GithubClient {
	return &githubClient{client: client}
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

func (cl *githubClient) GetCommits(
	ctx context.Context,
	owner string,
	name string,
	event int,
) ([]*github.RepositoryCommit, error) {
	commits, _, err := cl.client.PullRequests.ListCommits(ctx, owner, name, event, nil)

	return commits, err
}

func (cl *githubClient) Comment(
	ctx context.Context,
	owner string,
	name string,
	event int,
	comment *github.PullRequestComment,
) error {
	_, _, err := cl.client.PullRequests.CreateComment(ctx, owner, name, event, comment)

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

	cl.client.PullRequests.Edit(ctx, owner, name, event, &github.PullRequest{
		State: &state,
	})
	return nil
}