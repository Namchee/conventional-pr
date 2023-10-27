package mocks

import (
	"context"
	"errors"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

// GitHub's client mock. Used in testing
type githubClientMock struct{}

func (m *githubClientMock) GetPullRequest(
	_ context.Context,
	_ string,
	_ string,
	event int,
) (*github.PullRequest, error) {
	if event == 123 {
		return &github.PullRequest{}, nil
	}

	return nil, errors.New("not found")
}

func (m *githubClientMock) GetUser(ctx context.Context, name string) (*github.User, error) {
	if name == "foo" {
		return &github.User{Type: github.String(constants.BotUser)}, nil
	}

	return &github.User{
		Type: github.String(constants.NormalUser),
		ID:   github.Int64(123),
	}, nil
}

func (m *githubClientMock) GetComments(_ context.Context, _ string, _ string, number int) ([]*github.IssueComment, error) {
	if number == 1 {
		return nil, errors.New("error")
	}

	if number == 2 {
		return []*github.IssueComment{}, nil
	}

	return []*github.IssueComment{
		{
			ID:   github.Int64(3),
			Body: github.String("foo bar"),
			User: &github.User{
				ID: github.Int64(123),
			},
		},
	}, nil
}

func (m *githubClientMock) GetIssue(
	ctx context.Context,
	_ *entity.Meta,
	number int,
) (*github.Issue, error) {
	if number == 123 {
		return &github.Issue{}, nil
	}

	return nil, nil
}

func (m *githubClientMock) GetPermissionLevel(
	_ context.Context,
	_ *entity.Meta,
	user string,
) ([]string, error) {
	admin := constants.AdminUser
	writeOnly := "write"

	switch user {
	case "foo":
		return []string{admin}, nil
	case "bar":
		return []string{writeOnly}, nil
	default:
		return []string{}, errors.New("mock error")
	}
}

func (m *githubClientMock) GetCommits(
	ctx context.Context,
	_ *entity.Meta,
	event int,
) ([]*entity.Commit, error) {
	good := "feat(test): test something"
	bad := "this is bad"

	if event == 123 {
		return []*entity.Commit{
			{
				Message: good,
			},
		}, nil
	}

	if event == 69 {
		return []*entity.Commit{
			{
				Hash:    "e21b424",
				Message: bad,
			},
		}, nil
	}

	if event == 1 {
		return []*entity.Commit{}, errors.New("mock error")
	}

	return []*entity.Commit{}, nil
}

func (m *githubClientMock) CreateComment(
	_ context.Context,
	_ string,
	_ string,
	event int,
	_ *github.IssueComment,
) error {
	if event == 123 {
		return nil
	}

	return errors.New("Error")
}

func (m *githubClientMock) EditComment(
	_ context.Context,
	_ string,
	_ string,
	id int64,
	_ *github.IssueComment,
) error {
	if id == 123 {
		return nil
	}

	return errors.New("Error")
}

func (m *githubClientMock) Label(
	_ context.Context,
	_ string,
	_ string,
	event int,
	_ string,
) error {
	if event == 123 {
		return nil
	}

	return errors.New("Error")
}

func (m *githubClientMock) Close(
	_ context.Context,
	_ string,
	_ string,
	event int,
) error {
	if event == 123 {
		return nil
	}

	return errors.New("Error")
}

func (m *githubClientMock) GetIssueReferences(
	_ context.Context,
	_ *entity.Meta,

)

// NewGithubClientMock creates a GitHub client mock for testing purposes
func NewGithubClientMock() internal.GithubClient {
	return &githubClientMock{}
}
