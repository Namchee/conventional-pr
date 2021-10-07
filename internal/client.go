package internal

import (
	"context"

	"github.com/google/go-github/v32/github"
)

// GithubClient handles all interaction with Github's API
// Designed this way for easier software testing
type GithubClient interface {
	GetUser(context.Context, string) (*github.User, *github.Response, error)
}

type githubClient struct {
	client *github.Client
}

func NewGithubClient(client *github.Client) GithubClient {
	return &githubClient{client: client}
}

func (cl *githubClient) GetUser(ctx context.Context, name string) (*github.User, *github.Response, error) {
	return cl.client.Users.Get(ctx, name)
}
