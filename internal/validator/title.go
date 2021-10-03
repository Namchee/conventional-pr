package validator

import (
	"regexp"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type titleValidator struct {
	Name string
}

func NewTitleValidator() internal.Validator {
	return &titleValidator{
		Name: "Pull request has valid title",
	}
}

func (v *titleValidator) IsValid(pullRequest *github.PullRequest, config *entity.Config) error {
	title := pullRequest.GetTitle()

	pattern := regexp.MustCompile(config.Pattern)

	if !pattern.Match([]byte(title)) {
		return constants.ErrInvalidTitle
	}

	return nil
}