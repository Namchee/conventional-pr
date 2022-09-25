package validator

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type issueValidator struct {
	client internal.GithubClient
	config *entity.Configuration
	meta   *entity.Meta
	Name   string
}

// NewIssueValidator creates a new validator that validates issue resolution
func NewIssueValidator(
	client internal.GithubClient,
	config *entity.Configuration,
	meta *entity.Meta,
) internal.Validator {
	return &issueValidator{
		Name:   constants.IssueValidatorName,
		client: client,
		config: config,
		meta:   meta,
	}
}

func (v *issueValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if !v.config.Issue {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: false,
			Result: nil,
		}
	}

	issue := pullRequest.GetLinks().GetIssue()

	if issue != nil {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: true,
			Result: nil,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Active: true,
		Result: constants.ErrNoIssue,
	}
}
