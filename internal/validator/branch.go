package validator

import (
	"regexp"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func IsBranchValid(
	config *entity.Configuration,
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if config.BranchPattern == "" {
		return &entity.ValidationResult{
			Name:   constants.BranchValidatorName,
			Active: false,
			Result: nil,
		}
	}

	branch := pullRequest.Branch
	pattern := regexp.MustCompile(config.BranchPattern)

	if !pattern.Match([]byte(branch)) {
		return &entity.ValidationResult{
			Name:   constants.BranchValidatorName,
			Active: true,
			Result: constants.ErrInvalidBranch,
		}
	}

	return &entity.ValidationResult{
		Name:   constants.BranchValidatorName,
		Active: true,
		Result: nil,
	}
}
