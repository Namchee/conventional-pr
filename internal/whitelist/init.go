package whitelist

import (
	"sort"
	"sync"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/entity"
)

var (
	whitelists = []func(internal.GithubClient, *entity.Configuration) internal.Whitelist{
		NewBotWhitelist,
		NewDraftWhitelist,
		NewPermissionWhitelist,
		NewUsernameWhitelist,
	}
)

// WhitelistGroup is a collection of whitelisting process, integrated in one single function call
type WhitelistGroup struct {
	client internal.GithubClient
	config *entity.Configuration

	wg     *sync.WaitGroup
}

// NewWhitelistGroup creates a new WhitelistGroup
func NewWhitelistGroup(
	client internal.GithubClient,
	config *entity.Configuration,
	wg *sync.WaitGroup,
) *WhitelistGroup {
	return &WhitelistGroup{
		client: client,
		config: config,
		wg:     wg,
	}
}

func (w *WhitelistGroup) processWhitelist(
	whitelist internal.Whitelist,
	pullRequest *entity.PullRequest,
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

// Process the pull request with all available whitelists
func (w *WhitelistGroup) Process(
	pullRequest *entity.PullRequest,
) []*entity.WhitelistResult {
	channel := make(chan *entity.WhitelistResult, len(whitelists))
	w.wg.Add(len(whitelists))

	for _, wv := range whitelists {
		wl := wv(w.client, w.config)

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

		return results[i].Result
	})

	return results
}

// IsWhitelisted checks if a pull request is whitelisted or not from whitelist results
func IsWhitelisted(result []*entity.WhitelistResult) bool {
	for _, r := range result {
		if r.Active && r.Result {
			return true
		}
	}

	return false
}
