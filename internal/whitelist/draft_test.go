package whitelist

import (
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
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
				Result: false,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				Draft: &tc.args.draft,
			}
			config := &entity.Config{
				Draft: tc.args.config,
			}

			whitelister := NewDraftWhitelist(nil, config, nil)

			got := whitelister.IsWhitelisted(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
