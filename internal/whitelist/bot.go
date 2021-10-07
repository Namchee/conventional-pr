package whitelist

import (
	"context"
	"strings"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type botWhitelist struct {
	client internal.GithubClient
	config *entity.Config
	Name   string
}

func NewBotWhitelist(client internal.GithubClient, config *entity.Config) internal.Whitelist {
	return &botWhitelist{
		client: client,
		config: config,
		Name:   "Pull request is made by bot",
	}
}

func (w *botWhitelist) IsWhitelisted(pullRequest *github.PullRequest) bool {
	user, _, _ := w.client.GetUser(
		context.Background(),
		pullRequest.GetUser().GetLogin(),
	)

	return strings.ToLower(user.GetType()) == constants.BotUser &&
		w.config.Bot
}
