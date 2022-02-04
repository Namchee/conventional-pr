package whitelist

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/utils"
	"github.com/google/go-github/v32/github"
)

type usernameWhitelist struct {
	client internal.GithubClient
	config *entity.Config
	Name   string
}

// NewUsernameWhitelist creates a whitelist that bypasses checks for certain usernames
func NewUsernameWhitelist(client internal.GithubClient, config *entity.Config, _ *entity.Meta) internal.Whitelist {
	return &usernameWhitelist{
		client: client,
		config: config,
		Name:   constants.UsernameWhitelistName,
	}
}

func (w *usernameWhitelist) IsWhitelisted(pullRequest *github.PullRequest) *entity.WhitelistResult {
	if len(w.config.IgnoredUsers) == 0 {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	user := pullRequest.GetUser().GetLogin()

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: utils.ContainsString(w.config.IgnoredUsers, user),
	}
}
