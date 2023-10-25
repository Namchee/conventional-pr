package whitelist

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUsernameWhitelist_IsWhitelisted(t *testing.T) {
	type args struct {
		name   string
		config []string
	}
	tests := []struct {
		name string
		args args
		want *entity.WhitelistResult
	}{
		{
			name: "should be skipped if config is empty",
			args: args{
				name:   "foo",
				config: []string{},
			},
			want: &entity.WhitelistResult{
				Name:   constants.UsernameWhitelistName,
				Active: false,
				Result: false,
			},
		},
		{
			name: "should be checked if user is not on whitelist",
			args: args{
				name:   "foo",
				config: []string{"bar"},
			},
			want: &entity.WhitelistResult{
				Name:   constants.UsernameWhitelistName,
				Active: true,
				Result: false,
			},
		},
		{
			name: "should be skipped if user is on whitelist",
			args: args{
				name:   "bar",
				config: []string{"bar"},
			},
			want: &entity.WhitelistResult{
				Name:   constants.UsernameWhitelistName,
				Active: true,
				Result: true,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &entity.PullRequest{
				Author: entity.Actor{
					Login: tc.args.name,
				},
			}
			config := &entity.Configuration{
				IgnoredUsers: tc.args.config,
			}
			client := mocks.NewGithubClientMock()

			whitelister := NewUsernameWhitelist(client, config, nil)

			got := whitelister.IsWhitelisted(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
