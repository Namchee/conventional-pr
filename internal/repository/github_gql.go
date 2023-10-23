package repository

import (
	"context"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type githubGQLClient struct {
	client *githubv4.Client
}

// NewGithubGQLClient instatiates a new GraphQL GitHub client
func NewGithubGQLClient(config *entity.Configuration) internal.GithubClient {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)

	oauth := oauth2.NewClient(ctx, ts)
	github := githubv4.NewClient(oauth)

	if config.BaseURL != nil {
		github = githubv4.NewEnterpriseClient(config.BaseURL.String(), oauth)
	}

	return &githubGQLClient{client: github}
}