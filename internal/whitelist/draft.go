package whitelist

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
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
		Result: pullRequest.GetDraft(),
	}
}
