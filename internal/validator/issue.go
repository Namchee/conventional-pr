package validator

import (
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func IsIssueValid(config *entity.Configuration, pullRequest *entity.PullRequest) *entity.ValidationResult {
	if !config.Issue {
		return &entity.ValidationResult{
			Name:   constants.IssueValidatorName,
			Active: false,
			Result: nil,
		}
	}

	references := pullRequest.References
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
