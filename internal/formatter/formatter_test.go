package formatter

import (
	"errors"
	"testing"

	"github.com/Namchee/ethos/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestFormatResult(t *testing.T) {
	type args struct {
		whitelist  []*entity.WhitelistResult
		validation []*entity.ValidationResult
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "formats whitelisted pull request correctly",
			args: args{
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo bar",
						Result: true,
					},
				},
				validation: []*entity.ValidationResult{},
			},
			want: `## **Pull Request Validation Report**

*This comment is automatically generated by [ethos](https://github.com/Namchee/conventional-pr)*

### **Whitelist Report**

- ✅ foo bar

**Result**

Pull request matches with one (or more) enabled whitelist criteria. Pull request validation is skipped.`,
		},
		{
			name: "formats valid pull request correctly",
			args: args{
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo bar",
						Result: false,
					},
				},
				validation: []*entity.ValidationResult{
					{
						Name:   "bar baz",
						Result: nil,
					},
				},
			},
			want: `## **Pull Request Validation Report**

*This comment is automatically generated by [ethos](https://github.com/Namchee/conventional-pr)*

### **Whitelist Report**

- ❌ foo bar

**Result**

Pull request does not satisfy any enabled whitelist criteria. Pull request will be validated.

### **Validation Report**

- ✅ bar baz

**Result**

Pull request satisfies all enabled pull request rules.`,
		},
		{
			name: "format invalid pull request correctly",
			args: args{
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo bar",
						Result: false,
					},
				},
				validation: []*entity.ValidationResult{
					{
						Name:   "bar baz",
						Result: errors.New("testing"),
					},
				},
			},
			want: `## **Pull Request Validation Report**

*This comment is automatically generated by [ethos](https://github.com/Namchee/conventional-pr)*

### **Whitelist Report**

- ❌ foo bar

**Result**

Pull request does not satisfy any enabled whitelist criteria. Pull request will be validated.

### **Validation Report**

- ❌ bar baz

**Result**

Pull request is invalid.

**Reason**

- Testing`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatResult(tc.args.whitelist, tc.args.validation)

			assert.Equal(t, tc.want, got)
		})
	}
}
