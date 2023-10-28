package validator

import (
	"context"
	"sort"
	"sync"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/entity"
)

var (
	validators = []func(internal.GithubClient, *entity.Configuration) internal.Validator{
		NewTitleValidator,
		NewBodyValidator,
		NewBranchValidator,
		NewCommitValidator,
		NewIssueValidator,
		NewFileValidator,
	}
)

// ValidatorGroup is a collection of validation process, integrated in one function call
type ValidatorGroup struct {
	client internal.GithubClient
	config *entity.Configuration

	wg *sync.WaitGroup
}

// NewValidatorGroup creates a new ValidatorGroup
func NewValidatorGroup(
	client internal.GithubClient,
	config *entity.Configuration,
	wg *sync.WaitGroup,
) *ValidatorGroup {
	return &ValidatorGroup{
		config: config,
		client: client,
		wg:     wg,
	}
}

func (v *ValidatorGroup) processValidator(
	ctx context.Context,
	validator internal.Validator,
	pullRequest *entity.PullRequest,
	pool chan *entity.ValidationResult,
) {
	defer v.wg.Done()
	result := validator.IsValid(ctx, pullRequest)
	pool <- result
}

func (v *ValidatorGroup) cleanup(
	channel chan *entity.ValidationResult,
) {
	v.wg.Wait()
	close(channel)
}

// Process the pull request with all available validators
func (v *ValidatorGroup) Process(
	ctx context.Context,
	pullRequest *entity.PullRequest,
) []*entity.ValidationResult {
	channel := make(chan *entity.ValidationResult, len(validators))

	v.wg.Add(len(validators))

	for _, vv := range validators {
		va := vv(v.client, v.config)

		go v.processValidator(ctx, va, pullRequest, channel)
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

		return results[i].Result == nil
	})

	return results
}

// IsValid checks if a pull request is valid or not from validation results
func IsValid(result []*entity.ValidationResult) bool {
	for _, r := range result {
		if r.Active && r.Result != nil {
			return false
		}
	}

	return true
}
