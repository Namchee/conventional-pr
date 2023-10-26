package internal

import (
	"context"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// GithubClient handles all interaction with Github's API
// Designed this way for easier testing
type GithubClient interface {
	GetPullRequest(context.Context, *entity.Meta, int) (*entity.PullRequest, error)
	GetIssueReference(context.Context, *entity.Meta, int) ([]*entity.IssueReference, error)
	GetCommits(context.Context, *entity.Meta, int) ([]*entity.Commit, error)
	/*
	GetUser(context.Context, string) (*github.User, error)
	GetIssue(context.Context, string, string, int) (*github.Issue, error)
	GetComments(context.Context, *entity.Meta, int) ([]*github.IssueComment, error)
	GetPermissionLevel(context.Context, string, string, string) (*github.RepositoryPermissionLevel, error)
	GetCommits(context.Context, string, string, int) ([]*github.RepositoryCommit, error)
	CreateComment(context.Context, string, string, int, *github.IssueComment) error
	EditComment(context.Context, string, string, int64, *github.IssueComment) error
	Label(context.Context, string, string, int, string) error
	Close(context.Context, string, string, int) error
	*/
}

type githubClient struct {
	client *githubv4.Client
}

// NewGithubClient instatiates a new GitHub client API interface
func NewGithubClient(config *entity.Configuration) GithubClient {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)

	oauth := oauth2.NewClient(ctx, ts)
	github := githubv4.NewEnterpriseClient(config.BaseURL.String(), oauth)

	return &githubClient{client: github}
}

func (c *githubClient) GetPullRequest(
	ctx context.Context,
	meta *entity.Meta,
	number int,
) (*entity.PullRequest, error) {
	var query struct {
		Repository struct {
			PullRequest struct {
				Title string
				Body string
				HeadRefName string
				ChangedFiles int
				IsDraft bool
				Author struct {
					Type string `graphql:"typename: __typename"`
					Login string
				}
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(meta.Owner),
		"name": githubv4.String(meta.Name),
		"number": githubv4.Int(number),
	}

	err := c.client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	return &entity.PullRequest{
		Number: number,
		Title: query.Repository.PullRequest.Title,
		Body: query.Repository.PullRequest.Body,
		Branch: query.Repository.PullRequest.HeadRefName,
		IsDraft: query.Repository.PullRequest.IsDraft,
		Changes: query.Repository.PullRequest.ChangedFiles,
		Repository: *meta,
		Author: entity.Actor{
			Type: query.Repository.PullRequest.Author.Type,
			Login: query.Repository.PullRequest.Author.Login,
		},
	}, nil
}
