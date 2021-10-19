package validator

import (
	"regexp"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type branchValidator struct {
	Name   string
	client internal.GithubClient
	config *entity.Config
}

func NewBranchValidator(
	client internal.GithubClient,
	config *entity.Config,
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
			Result: nil,
		}
	}

	branch := pullRequest.GetHead().GetRef()

	pattern := regexp.MustCompile(v.config.BranchPattern)

	if !pattern.Match([]byte(branch)) {
		return &entity.ValidationResult{
			Name:   v.Name,
			Result: constants.ErrInvalidBranch,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Result: nil,
	}
}