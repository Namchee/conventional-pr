package internal

import (
	"context"

	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/formatter"
	"github.com/google/go-github/v32/github"
)

type GithubService struct {
	client GithubClient
	meta   *entity.Meta
}

func NewGithubService(
	client GithubClient,
	meta *entity.Meta,
) *GithubService {
	return &GithubService{
		client: client,
		meta:   meta,
	}
}

func (s *GithubService) WriteReport(
	pullRequest *github.PullRequest,
	whitelistResults []*entity.WhitelistResult,
	validationResults []*entity.ValidationResult,
) {
	report := formatter.FormatResult(whitelistResults, validationResults)

	ctx := context.Background()

	s.client.Comment(
		ctx,
		s.meta.Owner,
		s.meta.Name,
		pullRequest.GetNumber(),
		&github.PullRequestComment{
			Body: &report,
		},
	)
}
