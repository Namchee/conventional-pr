package whitelist

import (
	"context"
	"strings"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type permissionWhitelist struct {
	Name string

	client internal.GithubClient
	config *entity.Configuration
}

// NewPermissionWhitelist creates a whitelist that bypasses checks on pull request submitted by user with high privileges
func NewPermissionWhitelist(
	client internal.GithubClient,
	config *entity.Configuration,
) internal.Whitelist {
	return &permissionWhitelist{
		Name:   constants.PermissionWhitelistName,
		client: client,
		config: config,
	}
}

func (w *permissionWhitelist) IsWhitelisted(
	ctx context.Context,
	pullRequest *entity.PullRequest,
) *entity.WhitelistResult {
	if w.config.Strict {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	permissions, err := w.client.GetPermissions(
		ctx,
		&pullRequest.Repository,
		pullRequest.Author.Login,
	)

	if err != nil {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: true,
			Result: false,
		}
	}

	for _, permission := range permissions {
		if strings.ToLower(permission) == constants.AdminUser {
			return &entity.WhitelistResult{
				Name:   w.Name,
				Active: true,
				Result: true,
			}
		}
	}

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: false,
	}
}
