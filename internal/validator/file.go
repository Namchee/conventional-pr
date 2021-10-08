package validator

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type fileValidator struct {
	client internal.GithubClient
	config *entity.Config
	Name   string
}

func NewFileValidator(
	client internal.GithubClient,
	config *entity.Config,
	_ *entity.Meta,
) internal.Validator {
	return &fileValidator{
		client: client,
		config: config,
		Name:   "Pull request does not introduce too much changes",
	}
}

func (v *fileValidator) IsValid(pullRequest *github.PullRequest) error {
	if v.config.FileChanges <= 0 {
		return nil
	}

	if pullRequest.GetChangedFiles() <= v.config.FileChanges {
		return nil
	}

	return constants.ErrTooManyChanges
}
