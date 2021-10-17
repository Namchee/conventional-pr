package whitelist

import (
	"sort"
	"sync"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type WhitelistGroup struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	wg     *sync.WaitGroup
}

func NewWhitelistGroup(
	client internal.GithubClient,
	config *entity.Config,
	meta *entity.Meta,
	wg *sync.WaitGroup,
) *WhitelistGroup {
	return &WhitelistGroup{
		client: client,
		config: config,
		meta:   meta,
		wg:     wg,
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
	pullRequest *github.PullRequest,
) []*entity.WhitelistResult {
	bot := NewBotWhitelist(w.client, w.config, w.meta)
	draft := NewDraftWhitelist(w.client, w.config, w.meta)
	perms := NewPermissionWhitelist(w.client, w.config, w.meta)

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

func IsWhitelisted(result []*entity.WhitelistResult) bool {
	for _, r := range result {
		if r.Result {
			return true
		}
	}

	return false
}
