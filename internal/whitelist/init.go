package whitelist

import (
	"sync"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type WhitelistGroup struct {
	wg *sync.WaitGroup
}

func NewWhitelistGroup(wg *sync.WaitGroup) *WhitelistGroup {
	return &WhitelistGroup{
		wg: wg,
	}
}

func (w *WhitelistGroup) processWhitelist(
	whitelist internal.Whitelist,
	pullRequest *github.PullRequest,
	pool chan *entity.WhitelistResult,
) {
	go func() {
		defer w.wg.Done()

		w.wg.Add(1)
		result := whitelist.IsWhitelisted(pullRequest)
		pool <- result
	}()
}

func (w *WhitelistGroup) Process(
	client internal.GithubClient,
	config *entity.Config,
	meta *entity.Meta,
	pullRequest *github.PullRequest,
) []*entity.WhitelistResult {
	bot := NewBotWhitelist(client, config, meta)
	draft := NewDraftWhitelist(client, config, meta)
	perms := NewPermissionWhitelist(client, config, meta)

	resultPool := make(chan *entity.WhitelistResult, 3)

	go w.processWhitelist(bot, pullRequest, resultPool)
	go w.processWhitelist(draft, pullRequest, resultPool)
	go w.processWhitelist(perms, pullRequest, resultPool)

	w.wg.Wait()

	var results []*entity.WhitelistResult

	for result := range resultPool {
		results = append(results, result)
	}

	return results
}
