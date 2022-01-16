package whitelist

import (
	"context"
	"strings"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type botWhitelist struct {
	client internal.GithubClient
	config *entity.Config
	Name   string
}

// NewBotWhitelist creates a whitelist that bypasses checks on pull request submitted by bots
func NewBotWhitelist(
	client internal.GithubClient,
	config *entity.Config,
	_ *entity.Meta,
) internal.Whitelist {
	return &botWhitelist{
		Name:   constants.BotWhitelistName,
		client: client,
		config: config,
	}
}

func (w *botWhitelist) IsWhitelisted(pullRequest *github.PullRequest) *entity.WhitelistResult {
	if !w.config.Bot {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	user, _ := w.client.GetUser(
		context.Background(),
		pullRequest.GetUser().GetLogin(),
	)

	result := strings.ToLower(user.GetType()) == constants.BotUser &&
		w.config.Bot

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: result,
	}
}
