package whitelist

import (
	"context"
	"strings"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type permissionWhitelist struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	Name   string
}

func NewPermissionWhitelist(
	client internal.GithubClient,
	config *entity.Config,
	meta *entity.Meta,
) internal.Whitelist {
	return &permissionWhitelist{
		client: client,
		config: config,
		meta:   meta,
		Name:   "Pull request has high privileges",
	}
}

func (w *permissionWhitelist) IsWhitelisted(pullRequest *github.PullRequest) bool {
	ctx := context.Background()

	perms, _ := w.client.GetPermissionLevel(
		ctx,
		w.meta.Owner,
		w.meta.Name,
		pullRequest.GetUser().GetLogin(),
	)

	return strings.ToLower(perms.GetPermission()) == constants.AdminUser &&
		!w.config.Strict
}
