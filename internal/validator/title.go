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

func NewTitleValidator(_ internal.GithubClient, config *entity.Config) internal.Validator {
	return &titleValidator{
		Name:   "Pull request has valid title",
		config: config,
	}
}

func (v *titleValidator) IsValid(pullRequest *github.PullRequest) error {
	title := pullRequest.GetTitle()

	pattern := regexp.MustCompile(v.config.Pattern)

	if !pattern.Match([]byte(title)) {
		return constants.ErrInvalidTitle
	}

	return nil
}
