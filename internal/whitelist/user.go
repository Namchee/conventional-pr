package whitelist

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/utils"
)

type usernameWhitelist struct {
	client internal.GithubClient
	config *entity.Configuration
	Name   string
}

// NewUsernameWhitelist creates a whitelist that bypasses checks for certain usernames
func NewUsernameWhitelist(client internal.GithubClient, config *entity.Configuration, _ *entity.Meta) internal.Whitelist {
	return &usernameWhitelist{
		client: client,
		config: config,
		Name:   constants.UsernameWhitelistName,
	}
}

func (w *usernameWhitelist) IsWhitelisted(pullRequest *entity.PullRequest) *entity.WhitelistResult {
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
