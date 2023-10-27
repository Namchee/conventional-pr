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
	GetIssueReferences(context.Context, *entity.Meta, int) ([]*entity.IssueReference, error)
	GetCommits(context.Context, *entity.Meta, int) ([]*entity.Commit, error)
	GetPermissions(context.Context, *entity.Meta, string) ([]string, error)

	/*
		GetUser(context.Context, string) (*github.User, error)
		GetIssue(context.Context, string, string, int) (*github.Issue, error)
		GetComments(context.Context, *entity.Meta, int) ([]*github.IssueComment, error)
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
	prNumber int,
) (*entity.PullRequest, error) {
	var query struct {
		Repository struct {
			PullRequest struct {
				Title        string
				Body         string
				HeadRefName  string
				ChangedFiles int
				IsDraft      bool
				Author       struct {
					Type  string `graphql:"typename: __typename"`
					Login string
				}
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner":  githubv4.String(meta.Owner),
		"name":   githubv4.String(meta.Name),
		"number": githubv4.Int(prNumber),
	}

	err := c.client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	return &entity.PullRequest{
		Number:     prNumber,
		Title:      query.Repository.PullRequest.Title,
		Body:       query.Repository.PullRequest.Body,
		Branch:     query.Repository.PullRequest.HeadRefName,
		IsDraft:    query.Repository.PullRequest.IsDraft,
		Changes:    query.Repository.PullRequest.ChangedFiles,
		Repository: *meta,
		Author: entity.Actor{
			Type:  query.Repository.PullRequest.Author.Type,
			Login: query.Repository.PullRequest.Author.Login,
		},
	}, nil
}

func (c *githubClient) GetIssueReferences(
	ctx context.Context,
	meta *entity.Meta,
	prNumber int,
) ([]*entity.IssueReference, error) {
	var query struct {
		Repository struct {
			PullRequest struct {
				ClosingIssuesReferences struct {
					PageInfo struct {
						EndCursor   string
						HasNextPage bool
					}
					Nodes []struct {
						Number     int
						Repository struct {
							Owner struct {
								Login string
							}
							Name string
						}
					}
				} `graphql:"closingIssueReference(first: 100, after: $cursor)"`
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner":  githubv4.String(meta.Owner),
		"name":   githubv4.String(meta.Name),
		"number": githubv4.Int(prNumber),
	}

	var allReferences []*entity.IssueReference
	for {
		err := c.client.Query(ctx, &query, variables)
		if err != nil {
			return allReferences, err
		}

		for _, node := range query.Repository.PullRequest.ClosingIssuesReferences.Nodes {
			allReferences = append(allReferences, &entity.IssueReference{
				Number: node.Number,
				Meta: entity.Meta{
					Owner: node.Repository.Owner.Login,
					Name:  node.Repository.Name,
				},
			})
		}

		if !query.Repository.PullRequest.ClosingIssuesReferences.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = query.Repository.PullRequest.ClosingIssuesReferences.PageInfo.EndCursor
	}

	return allReferences, nil
}

func (c *githubClient) GetCommits(
	ctx context.Context,
	meta *entity.Meta,
	prNumber int,
) ([]*entity.Commit, error) {
	var query struct {
		Repository struct {
			PullRequest struct {
				Commits struct {
					PageInfo struct {
						EndCursor   string
						HasNextPage bool
					}
					Nodes []struct {
						Commit struct {
							Message string
							Oid     string
						}
					}
				} `graphql:"commits(first: 100, after: $cursor)"`
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner":  githubv4.String(meta.Owner),
		"name":   githubv4.String(meta.Name),
		"number": githubv4.Int(prNumber),
	}

	var allCommits []*entity.Commit
	for {
		err := c.client.Query(ctx, &query, variables)
		if err != nil {
			return allCommits, err
		}

		for _, node := range query.Repository.PullRequest.Commits.Nodes {
			allCommits = append(allCommits, &entity.Commit{
				Hash:    node.Commit.Oid,
				Message: node.Commit.Message,
			})
		}

		if !query.Repository.PullRequest.Commits.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = query.Repository.PullRequest.Commits.PageInfo.EndCursor
	}

	return allCommits, nil
}

func (c *githubClient) GetPermissions(
	ctx context.Context,
	meta *entity.Meta,
	login string,
) ([]string, error) {
	var query struct {
		Repository struct {
			Collaborators struct {
				Edges []struct {
					Permission string
				}
			} `graphql:"collaborators(login: $login)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(meta.Owner),
		"name":  githubv4.String(meta.Name),
		"login": githubv4.String(login),
	}

	var allPermissions []string

	err := c.client.Query(ctx, &query, variables)
	if err != nil {
		return allPermissions, err
	}

	for _, node := range query.Repository.Collaborators.Edges {
		allPermissions = append(allPermissions, node.Permission)
	}

	return allPermissions, nil
}
