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
		Name:   constants.FileValidatorName,
		client: client,
		config: config,
	}
}

func (v *fileValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if v.config.FileChanges == 0 {
		return &entity.ValidationResult{
			Name:   v.Name,
			Result: nil,
		}
	}

	if pullRequest.GetChangedFiles() <= v.config.FileChanges {
		return &entity.ValidationResult{
			Name:   v.Name,
			Result: nil,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Result: constants.ErrTooManyChanges,
	}
}
