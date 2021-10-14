package whitelist

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type draftWhitelist struct {
	config *entity.Config
	Name   string
}

func NewDraftWhitelist(
	_ internal.GithubClient,
	config *entity.Config,
	_ *entity.Meta,
) internal.Whitelist {
	return &draftWhitelist{
		Name:   "Pull request is a draft",
		config: config,
	}
}

func (w *draftWhitelist) IsWhitelisted(pullRequest *github.PullRequest) *entity.WhitelistResult {
	return &entity.WhitelistResult{
		Name:   w.Name,
		Result: pullRequest.GetDraft() && !w.config.Draft,
	}
}
