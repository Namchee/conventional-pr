package internal

import (
	"context"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

// GithubClient handles all interaction with Github's API
// Designed this way for easier testing
type GithubClient interface {
	GetPullRequest(context.Context, string, string, int) (*entity.PullRequest, error)
	GetUser(context.Context, string) (*github.User, error)
	GetIssue(context.Context, string, string, int) (*github.Issue, error)
	GetComments(context.Context, string, string, int) ([]*github.IssueComment, error)
	GetPermissionLevel(context.Context, string, string, string) (*github.RepositoryPermissionLevel, error)
	GetCommits(context.Context, string, string, int) ([]*github.RepositoryCommit, error)
	CreateComment(context.Context, string, string, int, *github.IssueComment) error
	EditComment(context.Context, string, string, int64, *github.IssueComment) error
	Label(context.Context, string, string, int, string) error
	Close(context.Context, string, string, int) error
}
