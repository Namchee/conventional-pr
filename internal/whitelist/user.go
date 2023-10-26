package whitelist

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/utils"
)

type usernameWhitelist struct {
	Name   string

	config *entity.Configuration
}

// NewUsernameWhitelist creates a whitelist that bypasses checks for certain usernames
func NewUsernameWhitelist(
	_ internal.GithubClient,
	config *entity.Configuration,
) internal.Whitelist {
	return &usernameWhitelist{
		Name:   constants.UsernameWhitelistName,
		config: config,
	}
}

func (w *usernameWhitelist) IsWhitelisted(
	pullRequest *entity.PullRequest,
) *entity.WhitelistResult {
	if len(w.config.IgnoredUsers) == 0 {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	user := pullRequest.Author.Login

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: utils.ContainsString(w.config.IgnoredUsers, user),
	}
}
