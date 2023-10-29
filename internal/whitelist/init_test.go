package whitelist

import (
	"context"
	"sync"
	"testing"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestWhitelistGroup(t *testing.T) {
	clientMock := mocks.NewGithubClientMock()

	pullRequest := &entity.PullRequest{
		Number: 123,
	}

	config := &entity.Configuration{}

	wg := sync.WaitGroup{}

	whitelistGroup := NewWhitelistGroup(
		clientMock,
		config,
		&wg,
	)

	got := whitelistGroup.Process(context.TODO(), pullRequest)

	assert.Equal(t, 4, len(got))
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
					Active: true,
					Result: false,
				},
				{
					Name:   "bar baz",
					Active: true,
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
					Active: true,
					Result: false,
				},
				{
					Name:   "bar baz",
					Active: true,
					Result: false,
				},
			},
			want: false,
		},
		{
			name: "should ignore whitelist if inactive",
			args: []*entity.WhitelistResult{
				{
					Name:   "foo bar",
					Active: false,
					Result: true,
				},
				{
					Name:   "bar baz",
					Active: true,
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
