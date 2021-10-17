package validator

import (
	"sort"
	"sync"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type ValidatorGroup struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	wg     *sync.WaitGroup
}

func NewValidatorGroup(
	client internal.GithubClient,
	config *entity.Config,
	meta *entity.Meta,
	wg *sync.WaitGroup,
) *ValidatorGroup {
	return &ValidatorGroup{
		client: client,
		config: config,
		meta:   meta,
		wg:     wg,
	}
}

func (v *ValidatorGroup) processValidator(
	validator internal.Validator,
	pullRequest *github.PullRequest,
	pool chan *entity.ValidationResult,
) {
	defer v.wg.Done()
	result := validator.IsValid(pullRequest)
	pool <- result
}

func (v *ValidatorGroup) cleanup(
	channel chan *entity.ValidationResult,
) {
	v.wg.Wait()
	close(channel)
}

func (w *ValidatorGroup) Process(
	pullRequest *github.PullRequest,
) []*entity.ValidationResult {
	title := NewTitleValidator(w.client, w.config, w.meta)
	body := NewBodyValidator(w.client, w.config, w.meta)
	file := NewFileValidator(w.client, w.config, w.meta)
	issue := NewIssueValidator(w.client, w.config, w.meta)
	commit := NewCommitValidator(w.client, w.config, w.meta)

	channel := make(chan *entity.ValidationResult, 5)

	w.wg.Add(5)
	go w.processValidator(title, pullRequest, channel)
	go w.processValidator(body, pullRequest, channel)
	go w.processValidator(file, pullRequest, channel)
	go w.processValidator(issue, pullRequest, channel)
	go w.processValidator(commit, pullRequest, channel)

	go w.cleanup(channel)

	var results []*entity.ValidationResult

	for result := range channel {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Result == results[j].Result {
			return results[i].Name < results[j].Name
		}

		if results[i].Result == nil {
			return false
		}

		return true
	})

	return results
}

func IsValid(result []*entity.ValidationResult) bool {
	for _, r := range result {
		if r.Result != nil {
			return false
		}
	}

	return true
}
