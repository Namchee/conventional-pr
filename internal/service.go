package internal

import (
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type GithubService struct {
	client GithubClient
	config *entity.Config
	meta   *entity.Meta
}

func NewGithubService(
	client GithubClient,
	config *entity.Config,
	meta *entity.Meta,
) *GithubService {
	return &GithubService{
		client: client,
		config: config,
		meta:   meta,
	}
}

func (s *GithubService) Report(
	pullRequest *github.PullRequest,
	whitelistResults []*entity.WhitelistResult,
	validatorResults []*entity.ValidatorResult,
) {
	report := constants.ReportHeader
}
