package mocks

import (
	"context"

	"github.com/Namchee/ethos/internal"
	"github.com/google/go-github/v32/github"
)

// GitHub's client mock. Used in testing
type githubClientMock struct{}

func (m *githubClientMock) GetUser(ctx context.Context, name string) (*github.User, *github.Response, error) {
	bot := "bot"
	user := "user"

	if name == "foo" {
		return &github.User{Type: &bot}, nil, nil
	}

	return &github.User{Type: &user}, nil, nil
}

func NewGithubClientMock() internal.GithubClient {
	return &githubClientMock{}
}
