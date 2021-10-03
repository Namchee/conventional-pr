package validator

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/utils"
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

	pattern := regexp.MustCompile(`^([\w]+)\(([\w\- _]+)\)?!?: [\w\s]+`)

	if !pattern.Match([]byte(title)) {
		return constants.ErrInvalidTitle
	}

	submatches := pattern.FindStringSubmatch(title)

	if len(config.AllowedTypes) > 1 &&
		!utils.ContainsString(config.AllowedTypes, submatches[1]) {
		errString := fmt.Sprintf(
			"Pull request title contains an illegal type `%s`.",
			submatches[1],
		)

		return errors.New(errString)
	}

	if len(config.AllowedScopes) > 1 &&
		!utils.ContainsString(config.AllowedScopes, submatches[2]) {
			errString := fmt.Sprintf(
				"Pull request title contains an illegal scope `%s`.",
				submatches[1],
			)
	
			return errors.New(errString)
	}

	return nil
}