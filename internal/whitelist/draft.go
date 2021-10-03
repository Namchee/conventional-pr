package whitelist

import (
	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type draftWhitelist struct {
	Name string
}

func NewDraftWhitelist() internal.Whitelist {
	return &draftWhitelist{
		Name: "Pull request is a draft",
	}
}

func (w *draftWhitelist) IsWhitelisted(pullRequest *github.PullRequest, config *entity.Config) bool {
	return pullRequest.GetDraft() && !config.Draft 
}