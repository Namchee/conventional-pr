package whitelist

import (
	"context"
	"fmt"
	"strings"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type botWhitelist struct {
	client internal.GithubClient
	config *entity.Configuration
	Name   string
}

// NewBotWhitelist creates a whitelist that bypasses checks on pull request submitted by bots
func NewBotWhitelist(
	client internal.GithubClient,
	config *entity.Configuration,
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

	user, err := w.client.GetUser(
		context.Background(),
		pullRequest.GetUser().GetLogin(),
	)

	fmt.Println(err)

	result := strings.ToLower(user.GetType()) == constants.BotUser &&
		w.config.Bot

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: result,
	}
}
