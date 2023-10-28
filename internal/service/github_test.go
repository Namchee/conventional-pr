package service

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGithubClient_WriteReport(t *testing.T) {
	type args struct {
		number  int
		results *entity.PullRequestResult
	}
	tests := []struct {
		name    string
		args    args
		config  *entity.Configuration
		wantErr bool
	}{
		{
			name: "new comment - should return error",
			args: args{
				number: 1,
				results: &entity.PullRequestResult{
					Whitelist: []*entity.WhitelistResult{
						{
							Name:   "foo",
							Result: true,
						},
					},
					Validation: []*entity.ValidationResult{},
				},
			},
			config:  &entity.Configuration{},
			wantErr: true,
		},
		{
			name: "new comment - success",
			args: args{
				number: 123,
				results: &entity.PullRequestResult{
					Whitelist: []*entity.WhitelistResult{
						{
							Name:   "foo",
							Result: true,
						},
					},
					Validation: []*entity.ValidationResult{},
				},
			},
			config:  &entity.Configuration{},
			wantErr: false,
		},
		{
			name: "edit - error cannot get comments",
			args: args{
				number: 1,
				results: &entity.PullRequestResult{
					Whitelist: []*entity.WhitelistResult{
						{
							Name:   "foo",
							Result: true,
						},
					},
					Validation: []*entity.ValidationResult{},
				},
			},
			config: &entity.Configuration{
				Edit: true,
			},
			wantErr: true,
		},
		{
			name: "edit - cannot find comment",
			args: args{
				number: 2,
				results: &entity.PullRequestResult{
					Whitelist: []*entity.WhitelistResult{
						{
							Name:   "foo",
							Result: true,
						},
					},
					Validation: []*entity.ValidationResult{},
				},
			},
			config: &entity.Configuration{
				Edit: true,
			},
			wantErr: true,
		},
		{
			name: "edit - success",
			args: args{
				number: 123,
				results: &entity.PullRequestResult{
					Whitelist: []*entity.WhitelistResult{
						{
							Name:   "foo",
							Result: true,
						},
					},
					Validation: []*entity.ValidationResult{},
				},
			},
			config: &entity.Configuration{
				Edit: true,
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pullRequest := &entity.PullRequest{
				Number: tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			meta := &entity.Meta{}

			service := NewGithubService(client, tc.config, meta)

			got := service.WriteReport(
				pullRequest,
				tc.args.results,
				mocks.ClockMock{}.Now(),
			)

			assert.Equal(t, tc.wantErr, got != nil)
		})
	}
}

func TestGithubClient_WriteTemplate(t *testing.T) {
	type args struct {
		number   int
		template string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return error",
			args: args{
				number:   1,
				template: "foo",
			},
			wantErr: true,
		},
		{
			name: "should do nothing if template is empty",
			args: args{
				number:   123,
				template: "",
			},
			wantErr: false,
		},
		{
			name: "should not return error",
			args: args{
				number:   123,
				template: "foo",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pullRequest := &entity.PullRequest{
				Number: tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			config := &entity.Configuration{
				Message: tc.args.template,
			}
			meta := &entity.Meta{}

			service := NewGithubService(client, config, meta)

			got := service.WriteMessage(pullRequest)

			assert.Equal(t, tc.wantErr, got != nil)
		})
	}
}

func TestGithubClient_AttachLabel(t *testing.T) {
	type args struct {
		number int
		label  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return error",
			args: args{
				number: 1,
				label:  "foo",
			},
			wantErr: true,
		},
		{
			name: "should do nothing if label is empty",
			args: args{
				number: 123,
				label:  "",
			},
			wantErr: false,
		},
		{
			name: "should not return error",
			args: args{
				number: 123,
				label:  "foo",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pullRequest := &entity.PullRequest{
				Number: tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			config := &entity.Configuration{
				Label: tc.args.label,
			}
			meta := &entity.Meta{}

			service := NewGithubService(client, config, meta)

			got := service.AttachLabel(pullRequest)

			assert.Equal(t, tc.wantErr, got != nil)
		})
	}
}

func TestGithubClient_ClosePullRequest(t *testing.T) {
	type args struct {
		number int
		close  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return error",
			args: args{
				number: 1,
				close:  true,
			},
			wantErr: true,
		},
		{
			name: "should do nothing if close is false",
			args: args{
				number: 123,
				close:  false,
			},
			wantErr: false,
		},
		{
			name: "should not return error",
			args: args{
				number: 123,
				close:  true,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pullRequest := &entity.PullRequest{
				Number: tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			config := &entity.Configuration{
				Close: tc.args.close,
			}
			meta := &entity.Meta{}

			service := NewGithubService(client, config, meta)

			got := service.ClosePullRequest(pullRequest)

			assert.Equal(t, tc.wantErr, got != nil)
		})
	}
}
