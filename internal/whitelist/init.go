package whitelist

import (
	"sort"
	"sync"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

var (
	whitelists = []func(internal.GithubClient, *entity.Config, *entity.Meta) internal.Whitelist{
		NewBotWhitelist,
		NewDraftWhitelist,
		NewPermissionWhitelist,
	}
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
	channel := make(chan *entity.WhitelistResult, len(whitelists))
	w.wg.Add(len(whitelists))

	for _, wv := range whitelists {
		wl := wv(w.client, w.config, w.meta)

		go w.processWhitelist(wl, pullRequest, channel)
	}

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
			return true
		}

		return false
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
