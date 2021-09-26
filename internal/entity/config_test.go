package entity

import (
	"os"
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name  string
		mocks map[string]string
		want  Config
	}{
		{
			name: "should read config correctly",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN":         "foo_bar",
				"INPUT_CHECK_DRAFT":          "false",
				"INPUT_LINK_ISSUE":           "true",
				"INPUT_IGNORE_BOT":           "false",
				"INPUT_MAXIMUM_FILE_CHANGES": "11",
			},
			want: Config{
				Token:       "foo_bar",
				Draft:       false,
				Issue:       true,
				Bot:         false,
				FileChanges: 11,
			},
		},
		{
			name: "should fallback to github-actions account",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN": "",
				"GITHUB_TOKEN":       "baz",
				"INPUT_CHECK_DRAFT":  "false",
				"INPUT_LINK_ISSUE":   "true",
				"INPUT_IGNORE_BOT":   "false",
			},
			want: Config{
				Token: "baz",
				Draft: false,
				Issue: true,
				Bot:   false,
			},
		},
		{
			name: "should be able to handle arrays",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN":  "foo_bar",
				"INPUT_ALLOWED_TYPES": "a, b, c",
			},
			want: Config{
				Token:        "foo_bar",
				AllowedTypes: []string{"a", "b", "c"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for key, val := range tc.mocks {
				os.Setenv(key, val)
				defer os.Unsetenv(key)
			}

			got := ReadConfig()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ReadConfig() = %v, want = %v", got, tc.want)
			}
		})
	}
}
