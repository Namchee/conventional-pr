package validator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type commitValidator struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	Name   string
}

func NewCommitValidator(
	client internal.GithubClient,
	config *entity.Config,
	meta *entity.Meta,
) internal.Validator {
	return &commitValidator{
		client: client,
		config: config,
		meta:   meta,
		Name:   "All commit(s) has valid messages",
	}
}

func (v *commitValidator) IsValid(pullRequest *github.PullRequest) error {
	if !v.config.Commits {
		return nil
	}

	ctx := context.Background()

	commits, _ := v.client.GetCommits(
		ctx,
		v.meta.Owner,
		v.meta.Name,
		pullRequest.GetNumber(),
	)

	pattern := regexp.MustCompile(v.config.Pattern)

	fmt.Println(commits)

	for _, commit := range commits {
		message := commit.Commit.GetMessage()

		if !pattern.Match([]byte(message)) {
			return fmt.Errorf(
				"commit %s does not have valid commit message", commit.GetSHA(),
			)
		}
	}

	return nil
}
