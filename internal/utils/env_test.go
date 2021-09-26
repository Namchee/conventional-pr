package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnvBool(t *testing.T) {
	tests := []struct {
		name      string
		mockValue string
		want      bool
	}{
		{
			name:      "should read true correctly",
			mockValue: "true",
			want:      true,
		},
		{
			name:      "should read false correctly",
			mockValue: "false",
			want:      false,
		},
		{
			name:      "should fallback to false",
			mockValue: "bar",
			want:      false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("TEST", tc.mockValue)
			defer os.Unsetenv("TEST")

			got := ReadEnvBool("TEST")

			assert.Equal(
				t,
				got,
				tc.want,
				fmt.Sprintf("ReadEnvBool() = %v, want = %v", got, tc.want),
			)
		})
	}
}

func TestReadEnvString(t *testing.T) {
	tests := []struct {
		name      string
		mockValue string
		want      string
	}{
		{
			name:      "should read correctly",
			mockValue: "foo",
			want:      "foo",
		},
		{
			name:      "should return zero value",
			mockValue: "",
			want:      "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.mockValue) > 0 {
				os.Setenv("TEST", tc.mockValue)
				defer os.Unsetenv("TEST")
			}

			got := ReadEnvString("TEST")

			assert.Equal(
				t,
				got,
				tc.want,
				fmt.Sprintf("ReadEnvString() = %v, want = %v", got, tc.want),
			)
		})
	}
}

func TestReadEnvInt(t *testing.T) {
	tests := []struct {
		name      string
		mockValue string
		want      int
	}{
		{
			name:      "should read correctly",
			mockValue: "123",
			want:      123,
		},
		{
			name:      "should return zero value",
			mockValue: "Hee Hoo",
			want:      0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("TEST", tc.mockValue)
			defer os.Unsetenv("TEST")

			got := ReadEnvInt("TEST")

			assert.Equal(
				t,
				got,
				tc.want,
				fmt.Sprintf("ReadEnvInt() = %v, want = %v", got, tc.want),
			)
		})
	}
}
