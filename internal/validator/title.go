package validator

import (
	"regexp"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type titleValidator struct {
	Name   string
	config *entity.Config
}

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

func (v *titleValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidatorResult {
	title := pullRequest.GetTitle()

	pattern := regexp.MustCompile(v.config.Pattern)

	if !pattern.Match([]byte(title)) {
		return &entity.ValidatorResult{
			Name:   v.Name,
			Result: constants.ErrInvalidTitle,
		}
	}

	return &entity.ValidatorResult{
		Name:   v.Name,
		Result: nil,
	}
}
