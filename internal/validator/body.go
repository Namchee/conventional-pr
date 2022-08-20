package validator

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type bodyValidator struct {
	Name   string
	config *entity.Configuration
}

// NewBodyValidator creates a new validator that validates if a pull request has a non-empty body
func NewBodyValidator(
	_ internal.GithubClient,
	config *entity.Configuration,
	_ *entity.Meta,
) internal.Validator {
	return &bodyValidator{
		Name:   constants.BodyValidatorName,
		config: config,
	}
}

func (v *bodyValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if !v.config.Body {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: false,
			Result: nil,
		}
	}

	body := pullRequest.GetBody()

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
