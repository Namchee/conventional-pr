package whitelist

import (
	"sync"
	"testing"

	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestWhitelistGroup(t *testing.T) {
	clientMock := mocks.NewGithubClientMock()

	prNum := 123
	pullRequest := &github.PullRequest{
		Number: &prNum,
	}

	config := &entity.Config{}
	meta := &entity.Meta{}

	wg := sync.WaitGroup{}

	whitelistGroup := NewWhitelistGroup(
		clientMock,
		config,
		meta,
		&wg,
	)

	got := whitelistGroup.Process(pullRequest)

	assert.Equal(t, 3, len(got))
}

func TestIsWhitelisted(t *testing.T) {
	tests := []struct {
		name string
		args []*entity.WhitelistResult
		want bool
	}{
		{
			name: "should return true",
			args: []*entity.WhitelistResult{
				{
					Name:   "foo bar",
					Result: false,
				},
				{
					Name:   "bar baz",
					Result: true,
				},
			},
			want: true,
		},
		{
			name: "should return false",
			args: []*entity.WhitelistResult{
				{
					Name:   "foo bar",
					Result: false,
				},
				{
					Name:   "bar baz",
					Result: false,
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsWhitelisted(tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}
