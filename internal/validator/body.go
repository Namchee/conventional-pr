package validator

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type bodyValidator struct {
	Name string
}

func NewBodyValidator(
	_ internal.GithubClient,
	_ *entity.Config,
	_ *entity.Meta,
) internal.Validator {
	return &bodyValidator{
		Name: constants.BodyValidatorName,
	}
}

func (v *bodyValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	body := pullRequest.GetBody()

	if body != "" {
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
