package validator

import (
	"sort"
	"sync"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

var (
	validators = []func(internal.GithubClient, *entity.Config, *entity.Meta) internal.Validator{
		NewTitleValidator,
		NewBodyValidator,
		NewBranchValidator,
		NewCommitValidator,
		NewIssueValidator,
		NewFileValidator,
	}
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

func (v *ValidatorGroup) Process(
	pullRequest *github.PullRequest,
) []*entity.ValidationResult {
	channel := make(chan *entity.ValidationResult, len(validators))

	v.wg.Add(len(validators))

	for _, vv := range validators {
		va := vv(v.client, v.config, v.meta)

		go v.processValidator(va, pullRequest, channel)
	}

	go v.cleanup(channel)

	var results []*entity.ValidationResult

	for result := range channel {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Result == results[j].Result {
			return results[i].Name < results[j].Name
		}

		if results[i].Result == nil {
			return true
		}

		return false
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
