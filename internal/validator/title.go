package validator

import (
	"regexp"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type titleValidator struct {
	Name   string
	config *entity.Config
}

// NewTitleValidator creates a new validator that validates a pull request title
func NewTitleValidator(
	_ internal.GithubClient,
	config *entity.Config,
	_ *entity.Meta,
) internal.Validator {
	return &titleValidator{
		Name:   constants.TitleValidatorName,
		config: config,
	}
}

func (v *titleValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if v.config.TitlePattern == "" {
		return &entity.ValidationResult{
			Name:   v.Name,
			Result: nil,
		}
	}

	title := pullRequest.GetTitle()

	pattern := regexp.MustCompile(v.config.TitlePattern)

	if !pattern.Match([]byte(title)) {
		return &entity.ValidationResult{
			Name:   v.Name,
			Result: constants.ErrInvalidTitle,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Result: nil,
	}
}
