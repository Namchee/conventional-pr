package whitelist

import (
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestBotWhitelist_IsWhitelisted(t *testing.T) {
	type args struct {
		name   string
		config bool
	}
	tests := []struct {
		name string
		args args
		want *entity.WhitelistResult
	}{
		{
			name: "should be skipped if is bot and bot = true",
			args: args{
				name:   "foo",
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Result: true,
			},
		},
		{
			name: "should be checked if is bot and bot = false",
			args: args{
				name:   "foo",
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Result: false,
			},
		},
		{
			name: "should be checked if is not bot and bot = true",
			args: args{
				name:   "bar",
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Result: false,
			},
		},
		{
			name: "should be checked if is not bot and bot = false",
			args: args{
				name:   "bar",
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Result: false,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			user := &github.User{
				Login: &tc.args.name,
			}
			pull := &github.PullRequest{
				User: user,
			}
			config := &entity.Config{
				Bot: tc.args.config,
			}
			client := mocks.NewGithubClientMock()

			whitelister := NewBotWhitelist(client, config, nil)

			got := whitelister.IsWhitelisted(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
