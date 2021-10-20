package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "should capitalize lower-case string",
			args: "foo bar",
			want: "Foo bar",
		},
		{
			name: "should not change anything",
			args: "Foo bar",
			want: "Foo bar",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Capitalize(tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}
