package entity

import (
	"os"
	"testing"
	"testing/fstest"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestReadEvent(t *testing.T) {
	type want struct {
		event *Event
		err   error
	}
	tests := []struct {
		name     string
		path     string
		mockFile []byte
		want     want
	}{
		{
			name:     "throw error when file cannot be read",
			path:     `://///`,
			mockFile: []byte(`{ "foo": "bar" }`),
			want: want{
				event: nil,
				err:   constants.ErrEventFileRead,
			},
		},
		{
			name:     "throw error when file cannot be parsed",
			path:     "test.json",
			mockFile: []byte(`{ foo: "bar" }`),
			want: want{
				event: nil,
				err:   constants.ErrEventFileParse,
			},
		},
		{
			name:     "should return correctly",
			path:     "test.json",
			mockFile: []byte(`{ "action": "opened", "number": 1 }`),
			want: want{
				event: &Event{
					Action: "opened",
					Number: 1,
				},
				err: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("GITHUB_EVENT_PATH", tc.path)
			defer os.Unsetenv("GITHUB_EVENT_PATH")

			mock := fstest.MapFS{
				tc.path: {
					Data: tc.mockFile,
				},
			}

			got, err := ReadEvent(mock)

			assert.Equal(t, tc.want.event, got)
			assert.Equal(t, tc.want.err, err)
		})
	}
}
