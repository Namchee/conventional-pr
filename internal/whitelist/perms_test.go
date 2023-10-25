package whitelist

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPermissionWhitelist_IsWhitelisted(t *testing.T) {
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
			name: "should be skipped if is high privilege and strict = false",
			args: args{
				name:   "foo",
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.PermissionWhitelistName,
				Active: true,
				Result: true,
			},
		},
		{
			name: "should be checked if is high privilege and strict = true",
			args: args{
				name:   "foo",
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.PermissionWhitelistName,
				Active: false,
				Result: false,
			},
		},
		{
			name: "should be checked if is low privilege and strict = true",
			args: args{
				name:   "bar",
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.PermissionWhitelistName,
				Active: false,
				Result: false,
			},
		},
		{
			name: "should be checked if is low privilege and strict = false",
			args: args{
				name:   "bar",
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.PermissionWhitelistName,
				Active: true,
				Result: false,
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
				Strict: tc.args.config,
			}
			meta := &entity.Meta{}
			client := mocks.NewGithubClientMock()

			whitelister := NewPermissionWhitelist(client, config, meta)

			got := whitelister.IsWhitelisted(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
