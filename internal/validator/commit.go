package validator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type commitValidator struct {
	Name string

	client internal.GithubClient
	config *entity.Configuration
}

// NewCommitValidator creates a new validator that will validate all commit messages in a pull request
func NewCommitValidator(
	client internal.GithubClient,
	config *entity.Configuration,
) internal.Validator {
	return &commitValidator{
		Name:   constants.CommitValidatorName,
		client: client,
		config: config,
	}
}

func (v *commitValidator) IsValid(
	ctx context.Context,
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if v.config.CommitPattern == "" {
		return &entity.ValidationResult{
			Name:   constants.CommitValidatorName,
			Active: false,
			Result: nil,
		}
	}

	commits, err := v.client.GetCommits(
		ctx,
		&pullRequest.Repository,
		pullRequest.Number,
	)
	if err != nil {
		return &entity.ValidationResult{
			Name:   constants.CommitValidatorName,
			Active: true,
			Result: nil,
		}
	}

	pattern := regexp.MustCompile(v.config.CommitPattern)

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
