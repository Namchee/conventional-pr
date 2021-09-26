package entity

import (
	"reflect"
	"testing"

	"github.com/Namchee/ethos/internal/constants"
)

func TestCreateMeta(t *testing.T) {
	type result struct {
		meta *Meta
		err  error
	}
	tests := []struct {
		name    string
		args    string
		want    result
		wantErr bool
	}{
		{
			name: "should be able to extract metadata",
			args: "foo/bar",
			want: result{
				meta: &Meta{
					Name:  "bar",
					Owner: "foo",
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "should throw an error",
			args: "fake_github_repository",
			want: result{
				meta: nil,
				err:  constants.ErrMalformedMetadata,
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CreateMeta(tc.args)

			if tc.wantErr && err == nil {
				t.Fatalf("CreateMeta() err = %v, wantErr = %v", err, tc.wantErr)
			}

			if !reflect.DeepEqual(got, tc.want.meta) {
				t.Fatalf("CreateMeta() got = %v, wantErr = %v", got, tc.want.meta)
			}
		})
	}
}
