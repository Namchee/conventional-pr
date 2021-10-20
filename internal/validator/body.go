package validator

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type bodyValidator struct {
	Name   string
	config *entity.Config
}

// NewBodyValidator creates a new validator that validates if a pull request has a non-empty body
func NewBodyValidator(
	_ internal.GithubClient,
	config *entity.Config,
	_ *entity.Meta,
) internal.Validator {
	return &bodyValidator{
		Name:   constants.BodyValidatorName,
		config: config,
	}
}

func (v *bodyValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	body := pullRequest.GetBody()

	if body != "" || !v.config.Body {
		return &entity.ValidationResult{
			Name:   v.Name,
			Result: nil,
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Result: constants.ErrNoBody,
	}
}
