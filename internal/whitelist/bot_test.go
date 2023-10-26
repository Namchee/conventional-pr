package whitelist

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestBotWhitelist_IsWhitelisted(t *testing.T) {
	type args struct {
		name   string
		userType string
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
				userType: "Bot",
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Active: true,
				Result: true,
			},
		},
		{
			name: "should be checked if is bot and bot = false",
			args: args{
				name:   "foo",
				userType: "Bot",
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Active: false,
				Result: false,
			},
		},
		{
			name: "should be checked if is not bot and bot = true",
			args: args{
				name:   "bar",
				userType: "User",
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Active: true,
				Result: false,
			},
		},
		{
			name: "should be checked if is not bot and bot = false",
			args: args{
				name:   "bar",
				userType: "User",
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.BotWhitelistName,
				Active: false,
				Result: false,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &entity.PullRequest{
				Author: entity.Actor{
					Login: tc.args.name,
					Type: tc.args.userType,
				},
			}
			config := &entity.Configuration{
				Bot: tc.args.config,
			}

			whitelister := NewBotWhitelist(nil, config)
			got := whitelister.IsWhitelisted(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
