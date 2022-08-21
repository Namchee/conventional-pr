package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/formatter"
	"github.com/google/go-github/v32/github"
)

// GithubService is a service that simplifies GitHub API interaction
type GithubService struct {
	client internal.GithubClient
	config *entity.Configuration
	meta   *entity.Meta
}

// NewGithubService creates a new GitHub service that simplify API interaction with functions which is actually needed
func NewGithubService(
	client internal.GithubClient,
	config *entity.Configuration,
	meta *entity.Meta,
) *GithubService {
	return &GithubService{
		client: client,
		config: config,
		meta:   meta,
	}
}

// WriteReport creates a new comment that contains conventional-pr workflow report in markdown format
func (s *GithubService) WriteReport(
	pullRequest *github.PullRequest,
	results *entity.PullRequestResult,
	time time.Time,
) error {
	ctx := context.Background()

	report := formatter.FormatResultToTables(
		results,
		time,
	)

	if s.config.Edit {
		err := s.editComment(
			ctx,
			pullRequest.GetNumber(),
			report,
		)

		if err != nil &&
			errors.Is(err, constants.ErrFirstComment) {
			return s.createComment(
				ctx,
				pullRequest.GetNumber(),
				report,
			)
		}

		return err
	}

	return s.createComment(
		ctx,
		pullRequest.GetNumber(),
		report,
	)
}

func (s *GithubService) createComment(
	ctx context.Context,
	number int,
	body string,
) error {
	return s.client.CreateComment(
		ctx,
		s.meta.Owner,
		s.meta.Name,
		number,
		&github.IssueComment{
			Body: github.String(body),
		},
	)
}

func (s *GithubService) editComment(
	ctx context.Context,
	prNumber int,
	body string,
) error {
	var prev *github.IssueComment

	comments, err := s.client.GetComments(
		ctx,
		s.meta.Owner,
		s.meta.Name,
		prNumber,
	)
	if err != nil {
		return err
	}

	fmt.Println("here")

	user, err := s.client.GetUser(ctx, "")
	if err != nil {
		return err
	}

	for _, comment := range comments {
		if comment.GetUser().GetID() == user.GetID() {
			prev = comment
			break
		}
	}

	if prev == nil {
		return constants.ErrFirstComment
	}

	return s.client.EditComment(
		ctx,
		s.meta.Owner,
		s.meta.Name,
		prev.GetID(),
		&github.IssueComment{
			Body: github.String(body),
		},
	)
}

// WriteTemplate creates a new comment that contains user-desired message
func (s *GithubService) WriteTemplate(
	pullRequest *github.PullRequest,
) error {
	if s.config.Template == "" {
		return nil
	}

	ctx := context.Background()

	return s.client.CreateComment(
		ctx,
		s.meta.Owner,
		s.meta.Name,
		pullRequest.GetNumber(),
		&github.IssueComment{
			Body: &s.config.Template,
		},
	)
}

// AttachLabel attachs label to invalid pull request
func (s *GithubService) AttachLabel(
	pullRequest *github.PullRequest,
) error {
	if s.config.Label == "" {
		return nil
	}

	ctx := context.Background()

	return s.client.Label(
		ctx,
		s.meta.Owner,
		s.meta.Name,
		pullRequest.GetNumber(),
		s.config.Label,
	)
}

// ClosePullRequest closes invalid pull request
func (s *GithubService) ClosePullRequest(
	pullRequest *github.PullRequest,
) error {
	if !s.config.Close {
		return nil
	}

	ctx := context.Background()

	return s.client.Close(
		ctx,
		s.meta.Owner,
		s.meta.Name,
		pullRequest.GetNumber(),
	)
}
