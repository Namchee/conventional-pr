package utils

import (
	"reflect"
	"testing"
)

func TestRemoveEmptyStringsFromSlice(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "should not remove anything",
			args: []string{"a", "b"},
			want: []string{"a", "b"},
		},
		{
			name: "should remove empty strings",
			args: []string{"a", "", "c"},
			want: []string{"a", "c"},
		},
		{
			name: "should not remove whitespaces",
			args: []string{" "},
			want: []string{" "},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RemoveEmptyStrings(tc.args)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("RemoveEmptyStrings() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		input []string
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{
				input: []string{"foo", "bar"},
				val: "bar",
			},
			want: true,
		},
		{
			name: "should return false",
			args: args{
				input: []string{"foo", "bar"},
				val: "baz",
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ContainsString(tc.args.input, tc.args.val)

			if got != tc.want {
				t.Fatalf("ContainsString() = %v, want = %v", got, tc.want)
			}
		})
	}
}
