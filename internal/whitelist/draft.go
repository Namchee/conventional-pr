package whitelist

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type draftWhitelist struct {
	config *entity.Config
	Name   string
}

// NewDraftWhitelist creates a whitelist that bypasses draft pull request checks
func NewDraftWhitelist(
	_ internal.GithubClient,
	config *entity.Config,
	_ *entity.Meta,
) internal.Whitelist {
	return &draftWhitelist{
		Name:   constants.DraftWhitelistName,
		config: config,
	}
}

func (w *draftWhitelist) IsWhitelisted(pullRequest *github.PullRequest) *entity.WhitelistResult {
	return &entity.WhitelistResult{
		Name:   w.Name,
		Result: pullRequest.GetDraft() && w.config.Draft,
	}
}
