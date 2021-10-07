package mocks

import (
	"context"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/google/go-github/v32/github"
)

// GitHub's client mock. Used in testing
type githubClientMock struct{}

func (m *githubClientMock) GetUser(ctx context.Context, name string) (*github.User, error) {
	bot := constants.BotUser
	user := "user"

	if name == "foo" {
		return &github.User{Type: &bot}, nil
	}

	return &github.User{Type: &user}, nil
}

func (m *githubClientMock) GetPermissionLevel(
	ctx context.Context,
	owner string,
	name string,
	user string,
) (*github.RepositoryPermissionLevel, error) {
	admin := constants.AdminUser
	writeOnly := "write"

	if user == "foo" {
		return &github.RepositoryPermissionLevel{
			Permission: &admin,
		}, nil
	}

	return &github.RepositoryPermissionLevel{
		Permission: &writeOnly,
	}, nil
}

func NewGithubClientMock() internal.GithubClient {
	return &githubClientMock{}
}
