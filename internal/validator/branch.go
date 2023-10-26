package validator

import (
	"regexp"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type branchValidator struct {
	Name   string
	config *entity.Configuration
}

// NewBranchValidator creates a validator that validates pull request branch name
func NewBranchValidator(
	config *entity.Configuration,
) internal.Validator {
	return &branchValidator{
		Name:   constants.BranchValidatorName,
		config: config,
	}
}

func (v *branchValidator) IsValid(
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if v.config.BranchPattern == "" {
		return &entity.ValidationResult{
			Name:   constants.BranchValidatorName,
			Active: false,
			Result: nil,
		}
	}

	branch := pullRequest.Branch
	pattern := regexp.MustCompile(v.config.BranchPattern)

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
