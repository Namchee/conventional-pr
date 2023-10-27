package validator

import (
	"context"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type issueValidator struct {
	Name string

	client internal.GithubClient
	config *entity.Configuration
}

// NewIssueValidator creates a new validator that validates issue resolution
func NewIssueValidator(
	client internal.GithubClient,
	config *entity.Configuration,
) internal.Validator {
	return &issueValidator{
		Name:   constants.IssueValidatorName,
		client: client,
		config: config,
	}
}

func (v *issueValidator) IsValid(
	ctx context.Context,
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if !v.config.Issue {
		return &entity.ValidationResult{
			Name:   constants.IssueValidatorName,
			Active: false,
			Result: nil,
		}
	}

	references, err := v.client.GetIssueReferences(
		ctx,
		&pullRequest.Repository,
		pullRequest.Number,
	)
	if err != nil {
		return &entity.ValidationResult{
			Name:   constants.IssueValidatorName,
			Active: true,
			Result: nil,
		}
	}

	for _, reference := range references {
		repo := reference.Owner + "/" + reference.Name
		meta := pullRequest.Repository.Owner + "/" + pullRequest.Repository.Name

		if repo == meta {
			return &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: nil,
			}
		}
	}

	return &entity.ValidationResult{
		Name:   constants.IssueValidatorName,
		Active: true,
		Result: constants.ErrNoIssue,
	}
}
