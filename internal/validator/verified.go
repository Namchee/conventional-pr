package validator

import (
	"context"
	"fmt"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type verifiedValidator struct {
	client internal.GithubClient
	config *entity.Configuration
	meta   *entity.Meta
	Name   string
}

// NewVerifiedValidator creates a new validator that will validate all commit messages in a pull request
func NewVerifiedValidator(
	client internal.GithubClient,
	config *entity.Configuration,
	meta *entity.Meta,
) internal.Validator {
	return &verifiedValidator{
		Name:   constants.VerifiedCommitsValidatorName,
		client: client,
		config: config,
		meta:   meta,
	}
}

func (v *verifiedValidator) IsValid(
	pullRequest *github.PullRequest,
) *entity.ValidationResult {
	if !v.config.Verified {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: false,
			Result: nil,
		}
	}

	ctx := context.Background()

	commits, _ := v.client.GetCommits(
		ctx,
		v.meta.Owner,
		v.meta.Name,
		pullRequest.GetNumber(),
	)

	for _, commit := range commits {
		isVerified := commit.Commit.GetVerification().GetVerified()

		if !isVerified {
			return &entity.ValidationResult{
				Name:   v.Name,
				Active: true,
				Result: fmt.Errorf(
					"commit %s is not a verified commit", commit.GetSHA(),
				),
			}
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Active: true,
		Result: nil,
	}
}
