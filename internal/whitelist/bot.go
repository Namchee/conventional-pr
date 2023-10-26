package whitelist

import (
	"strings"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type botWhitelist struct {
	Name   string

	config *entity.Configuration
}

// NewBotWhitelist creates a whitelist that bypasses checks on pull request submitted by bots
func NewBotWhitelist(
	_ internal.GithubClient,
	config *entity.Configuration,
) internal.Whitelist {
	return &botWhitelist{
		Name:   constants.BotWhitelistName,
		config: config,
	}
}

func (w *botWhitelist) IsWhitelisted(
	pullRequest *entity.PullRequest,
) *entity.WhitelistResult {
	if !w.config.Bot {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	result := strings.ToLower(pullRequest.Author.Type) == constants.BotUser &&
		w.config.Bot

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: result,
	}
}	
