package validator

import (
	"context"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type bodyValidator struct {
	Name   string
	config *entity.Configuration
}

// NewBodyValidator creates a new validator that validates if a pull request has a non-empty body
func NewBodyValidator(
	_ internal.GithubClient,
	config *entity.Configuration,
) internal.Validator {
	return &bodyValidator{
		Name:   constants.BodyValidatorName,
		config: config,
	}
}

func (v *bodyValidator) IsValid(
	_ context.Context,
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if !v.config.Body {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: false,
			Result: nil,
		}
	}

	body := pullRequest.Body

	if body != "" {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: true,
			Result: nil,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Active: true,
		Result: constants.ErrNoBody,
	}
}
