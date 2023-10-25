package validator

import (
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func IsFileValid(config *entity.Configuration, pullRequest *entity.PullRequest) *entity.ValidationResult {
	if config.FileChanges == 0 {
		return &entity.ValidationResult{
			Name:   constants.FileValidatorName,
			Active: false,
			Result: nil,
		}
	}

	if pullRequest.Changes <= config.FileChanges {
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
