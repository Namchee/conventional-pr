package service

import (
	"testing"

	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestGithubClient_WriteReport(t *testing.T) {
	type args struct {
		number     int
		whitelist  []*entity.WhitelistResult
		validation []*entity.ValidationResult
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
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo",
						Result: true,
					},
				},
				validation: []*entity.ValidationResult{},
			},
			wantErr: true,
		},
		{
			name: "should not return error",
			args: args{
				number: 123,
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo",
						Result: true,
					},
				},
				validation: []*entity.ValidationResult{},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pullRequest := &github.PullRequest{
				Number: &tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			config := &entity.Config{}
			meta := &entity.Meta{}

			service := NewGithubService(client, config, meta)

			got := service.WriteReport(pullRequest, tc.args.whitelist, tc.args.validation)

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
			pullRequest := &github.PullRequest{
				Number: &tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			config := &entity.Config{
				Template: tc.args.template,
			}
			meta := &entity.Meta{}

			service := NewGithubService(client, config, meta)

			got := service.WriteTemplate(pullRequest)

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
			pullRequest := &github.PullRequest{
				Number: &tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			config := &entity.Config{
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
			pullRequest := &github.PullRequest{
				Number: &tc.args.number,
			}

			client := mocks.NewGithubClientMock()

			config := &entity.Config{
				Close: tc.args.close,
			}
			meta := &entity.Meta{}

			service := NewGithubService(client, config, meta)

			got := service.ClosePullRequest(pullRequest)

			assert.Equal(t, tc.wantErr, got != nil)
		})
	}
}
