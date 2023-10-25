package whitelist

import (
	"context"
	"strings"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type permissionWhitelist struct {
	client internal.GithubClient
	config *entity.Configuration
	meta   *entity.Meta
	Name   string
}

// NewPermissionWhitelist creates a whitelist that bypasses checks on pull request submitted by user with high privileges
func NewPermissionWhitelist(
	client internal.GithubClient,
	config *entity.Configuration,
	meta *entity.Meta,
) internal.Whitelist {
	return &permissionWhitelist{
		Name:   constants.PermissionWhitelistName,
		client: client,
		config: config,
		meta:   meta,
	}
}

func (w *permissionWhitelist) IsWhitelisted(pullRequest *entity.PullRequest) *entity.WhitelistResult {
	if w.config.Strict {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	ctx := context.Background()

	perms, _ := w.client.GetPermissionLevel(
		ctx,
		w.meta.Owner,
		w.meta.Name,
		pullRequest.Author.Login,
	)

	result := strings.ToLower(perms.GetPermission()) == constants.AdminUser

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: result,
	}
}
