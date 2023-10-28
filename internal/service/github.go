package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/formatter"
)

var (
	conventionalPrReportPattern = regexp.MustCompile("([Conventional PR])")
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
	pullRequest *entity.PullRequest,
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
			pullRequest.Number,
			report,
		)

		if err != nil &&
			errors.Is(err, constants.ErrFirstComment) {
			return s.createComment(
				ctx,
				pullRequest.Number,
				report,
			)
		}

		return err
	}

	return s.createComment(
		ctx,
		pullRequest.Number,
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
		s.meta,
		number,
		body,
	)
}

func (s *GithubService) editComment(
	ctx context.Context,
	prNumber int,
	body string,
) error {
	var prev *entity.Comment

	comments, err := s.client.GetComments(
		ctx,
		s.meta,
		prNumber,
	)
	if err != nil {
		return err
	}

	user, err := s.client.GetSelf(ctx)
	if err != nil {
		return err
	}

	for _, comment := range comments {
		isSameUser := comment.Author.Login == user.Login
		isReport := conventionalPrReportPattern.FindStringIndex(comment.Body)
		if isSameUser && len(isReport) > 0 {
			prev = comment
			break
		}
	}

	if prev == nil {
		return constants.ErrFirstComment
	}

	return s.client.EditComment(
		ctx,
		s.meta,
		prev.ID,
		body,
	)
}

// WriteMessage creates a new comment that contains user-desired message
func (s *GithubService) WriteMessage(
	pullRequest *entity.PullRequest,
) error {
	if s.config.Message == "" {
		return nil
	}

	ctx := context.Background()

	return s.client.CreateComment(
		ctx,
		s.meta,
		pullRequest.Number,
		s.config.Message,
	)
}

// AttachLabel attachs label to invalid pull request
func (s *GithubService) AttachLabel(
	pullRequest *entity.PullRequest,
) error {
	if s.config.Label == "" {
		return nil
	}

	ctx := context.Background()

	return s.client.Label(
		ctx,
		s.meta,
		pullRequest.Number,
		s.config.Label,
	)
}

// ClosePullRequest closes invalid pull request
func (s *GithubService) ClosePullRequest(
	pullRequest *entity.PullRequest,
) error {
	if !s.config.Close {
		return nil
	}

	ctx := context.Background()

	return s.client.Close(
		ctx,
		&pullRequest.Repository,
		pullRequest.Number,
	)
}
