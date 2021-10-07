package whitelist

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type botWhitelist struct {
	client *github.Client
	config *entity.Config
	Name   string
}

func NewBotWhitelist(client *github.Client) internal.Whitelist {
	return &botWhitelist{
		client: client,
		Name:   "Pull request is made by bot",
	}
}

func (w *botWhitelist) IsWhitelisted(pullRequest *github.PullRequest) bool {
	return true
}
