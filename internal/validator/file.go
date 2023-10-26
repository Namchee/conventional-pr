package validator

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type fileValidator struct {
	Name   string

	config *entity.Configuration
}

// NewFileValidator creates a new validator that validates if a pull request introduces too many file changes
func NewFileValidator(
	config *entity.Configuration,
) internal.Validator {
	return &fileValidator{
		Name:   constants.FileValidatorName,
		config: config,
	}
}

func (v *fileValidator) IsValid(pullRequest *entity.PullRequest) *entity.ValidationResult {
	if v.config.FileChanges == 0 {
		return &entity.ValidationResult{
			Name:   constants.FileValidatorName,
			Active: false,
			Result: nil,
		}
	}

	if pullRequest.Changes <= v.config.FileChanges {
		return &entity.ValidationResult{
			Name:   constants.FileValidatorName,
			Active: true,
			Result: nil,
		}
	}

	return &entity.ValidationResult{
		Name:   constants.FileValidatorName,
		Active: true,
		Result: constants.ErrTooManyChanges,
	}
}
