package whitelist

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type draftWhitelist struct {
	config *entity.Config
	Name   string
}

func NewDraftWhitelist(_ *github.Client, config *entity.Config) internal.Whitelist {
	return &draftWhitelist{
		Name:   "Pull request is a draft",
		config: config,
	}
}

func (w *draftWhitelist) IsWhitelisted(pullRequest *github.PullRequest) bool {
	return pullRequest.GetDraft() && !w.config.Draft
}
