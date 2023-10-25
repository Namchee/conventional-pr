package whitelist

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type draftWhitelist struct {
	config *entity.Configuration
	Name   string
}

// NewDraftWhitelist creates a whitelist that bypasses draft pull request checks
func NewDraftWhitelist(
	_ internal.GithubClient,
	config *entity.Configuration,
	_ *entity.Meta,
) internal.Whitelist {
	return &draftWhitelist{
		Name:   constants.DraftWhitelistName,
		config: config,
	}
}

func (w *draftWhitelist) IsWhitelisted(pullRequest *entity.PullRequest) *entity.WhitelistResult {
	if !w.config.Draft {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: pullRequest.IsDraft,
	}
}
