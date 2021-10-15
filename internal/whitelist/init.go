package whitelist

import (
	"sort"
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
	defer w.wg.Done()
	result := whitelist.IsWhitelisted(pullRequest)
	pool <- result
}

func (w *WhitelistGroup) cleanup(
	channel chan *entity.WhitelistResult,
) {
	w.wg.Wait()
	close(channel)
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

	channel := make(chan *entity.WhitelistResult, 3)

	w.wg.Add(3)
	go w.processWhitelist(bot, pullRequest, channel)
	go w.processWhitelist(draft, pullRequest, channel)
	go w.processWhitelist(perms, pullRequest, channel)

	go w.cleanup(channel)

	var results []*entity.WhitelistResult

	for result := range channel {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Result == results[j].Result {
			return results[i].Name < results[j].Name
		}

		if results[i].Result {
			return false
		}

		return true
	})

	return results
}
