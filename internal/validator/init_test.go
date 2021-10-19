package validator

import (
	"errors"
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

	assert.Equal(t, 6, len(got))
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		args []*entity.ValidationResult
		want bool
	}{
		{
			name: "should return true",
			args: []*entity.ValidationResult{
				{
					Name:   "foo bar",
					Result: nil,
				},
				{
					Name:   "bar baz",
					Result: nil,
				},
			},
			want: true,
		},
		{
			name: "should return false",
			args: []*entity.ValidationResult{
				{
					Name:   "foo bar",
					Result: nil,
				},
				{
					Name:   "bar baz",
					Result: errors.New("testing"),
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}
