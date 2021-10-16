package validator

import (
	"sync"
	"testing"

	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestValidatorGroup(t *testing.T) {
	clientMock := mocks.NewGithubClientMock()

	prNum := 123
	pullRequest := &github.PullRequest{
		Number: &prNum,
	}

	config := &entity.Config{}
	meta := &entity.Meta{}

	wg := sync.WaitGroup{}

	validatorGroup := NewValidatorGroup(
		clientMock,
		config,
		meta,
		&wg,
	)

	got := validatorGroup.Process(pullRequest)

	assert.Equal(t, 4, len(got))
}
