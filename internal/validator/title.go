package validator

import (
	"context"
	"regexp"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type titleValidator struct {
	Name string

	config *entity.Configuration
}

// NewTitleValidator creates a new validator that validates a pull request title
func NewTitleValidator(
	_ internal.GithubClient,
	config *entity.Configuration,
) internal.Validator {
	return &titleValidator{
		Name: constants.TitleValidatorName,

		config: config,
	}
}

func (v *titleValidator) IsValid(
	_ context.Context,
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if v.config.TitlePattern == "" {
		return &entity.ValidationResult{
			Name:   constants.TitleValidatorName,
			Active: false,
			Result: nil,
		}
	}

	title := pullRequest.Title
	pattern := regexp.MustCompile(v.config.TitlePattern)

	if !pattern.Match([]byte(title)) {
		return &entity.ValidationResult{
			Name:   constants.TitleValidatorName,
			Active: true,
			Result: constants.ErrInvalidTitle,
		}
	}

	return &entity.ValidationResult{
		Name:   constants.TitleValidatorName,
		Active: true,
		Result: nil,
	}
}
