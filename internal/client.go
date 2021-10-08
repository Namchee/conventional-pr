package internal

import (
	"context"

	"github.com/google/go-github/v32/github"
)

// GithubClient handles all interaction with Github's API
// Designed this way for easier software testing
type GithubClient interface {
	GetUser(context.Context, string) (*github.User, error)
	GetIssue(context.Context, string, string, int) (*github.Issue, error)
	GetPermissionLevel(context.Context, string, string, string) (*github.RepositoryPermissionLevel, error)
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
