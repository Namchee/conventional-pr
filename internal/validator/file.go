package validator

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type fileValidator struct {
	client internal.GithubClient
	config *entity.Configuration
	Name   string
}

// NewFileValidator creates a new validator that validates if a pull request introduces too many file changes
func NewFileValidator(
	client internal.GithubClient,
	config *entity.Configuration,
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
			Active: false,
			Result: nil,
		}
	}

	if pullRequest.GetChangedFiles() <= v.config.FileChanges {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: true,
			Result: nil,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Active: true,
		Result: constants.ErrTooManyChanges,
	}
}
