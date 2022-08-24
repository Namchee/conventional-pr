package validator

import (
	"regexp"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type branchValidator struct {
	Name   string
	client internal.GithubClient
	config *entity.Configuration
}

// NewBranchValidator creates a validator that validates pull request branch name
func NewBranchValidator(
	client internal.GithubClient,
	config *entity.Configuration,
	_ *entity.Meta,
) internal.Validator {
	return &branchValidator{
		Name:   constants.BranchValidatorName,
		client: client,
		config: config,
	}
}

func (v *branchValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if v.config.BranchPattern == "" {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: false,
			Result: nil,
		}
	}

	branch := pullRequest.GetHead().GetRef()

	pattern := regexp.MustCompile(v.config.BranchPattern)

	if !pattern.Match([]byte(branch)) {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: true,
			Result: constants.ErrInvalidBranch,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Active: true,
		Result: nil,
	}
}
