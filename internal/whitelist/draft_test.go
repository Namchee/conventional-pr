package whitelist

import (
	"context"
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestDraftWhitelist_IsWhitelisted(t *testing.T) {
	type args struct {
		draft  bool
		config bool
	}
	tests := []struct {
		name string
		args args
		want *entity.WhitelistResult
	}{
		{
			name: "should be skipped if config = false, draft = true",
			args: args{
				draft:  true,
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.DraftWhitelistName,
				Active: true,
				Result: true,
			},
		},
		{
			name: "should be checked if config = false, draft = true",
			args: args{
				draft:  true,
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.DraftWhitelistName,
				Active: false,
				Result: false,
			},
		},
		{
			name: "should be checked if config = false, draft = false",
			args: args{
				draft:  false,
				config: false,
			},
			want: &entity.WhitelistResult{
				Name:   constants.DraftWhitelistName,
				Active: false,
				Result: false,
			},
		},
		{
			name: "should be checked if config = true, draft = false",
			args: args{
				draft:  false,
				config: true,
			},
			want: &entity.WhitelistResult{
				Name:   constants.DraftWhitelistName,
				Active: true,
				Result: false,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &entity.PullRequest{
				IsDraft: tc.args.draft,
			}
			config := &entity.Configuration{
				Draft: tc.args.config,
			}

			whitelister := NewDraftWhitelist(nil, config)
			got := whitelister.IsWhitelisted(context.TODO(), pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
