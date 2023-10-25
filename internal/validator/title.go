package validator

import (
	"regexp"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func IsTitleValid(
	config *entity.Configuration,
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if config.TitlePattern == "" {
		return &entity.ValidationResult{
			Name:   constants.TitleValidatorName,
			Active: false,
			Result: nil,
		}
	}

	title := pullRequest.Title

	pattern := regexp.MustCompile(config.TitlePattern)

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
