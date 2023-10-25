package validator

import (
	"fmt"
	"regexp"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func IsCommitValid(config *entity.Configuration, pullRequest *entity.PullRequest) *entity.ValidationResult {
	if config.CommitPattern == "" {
		return &entity.ValidationResult{
			Name:   constants.CommitValidatorName,
			Active: false,
			Result: nil,
		}
	}

	commits := pullRequest.Commits
	pattern := regexp.MustCompile(config.CommitPattern)

	for _, commit := range commits {
		message := commit.Message

		if !pattern.Match([]byte(message)) {
			return &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Active: true,
				Result: fmt.Errorf(
					"commit %s does not have valid commit message", commit.Hash,
				),
			}
		}
	}

	return &entity.ValidationResult{
		Name:   constants.CommitValidatorName,
		Active: true,
		Result: nil,
	}
}
