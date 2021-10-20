package formatter

import (
	"fmt"
	"strings"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/utils"
)

func formatWhitelistResult(
	whitelistResults []*entity.WhitelistResult,
) string {
	flag := false
	var resultList []string

	for _, result := range whitelistResults {
		wResult := constants.FailEmoji

		if result.Result {
			flag = true
			wResult = constants.PassEmoji
		}

		wResult = fmt.Sprintf("- %s %s", wResult, result.Name)
		resultList = append(resultList, wResult)
	}

	report := constants.WhitelistFail

	if flag {
		report = constants.WhitelistPass
	}

	resultReport := fmt.Sprintf(constants.ResultTemplate, report)

	return fmt.Sprintf(constants.WhitelistTemplate, strings.Join(resultList, "\n"), resultReport)
}

func formatValidationResult(
	validationResults []*entity.ValidationResult,
) string {
	var fails []error
	var resultList []string

	for _, result := range validationResults {
		vResult := constants.PassEmoji

		if result.Result != nil {
			fails = append(fails, result.Result)
			vResult = constants.FailEmoji
		}

		vResult = fmt.Sprintf("- %s %s", vResult, result.Name)
		resultList = append(resultList, vResult)
	}

	var reason string
	report := constants.ValidationPass

	if len(fails) > 0 {
		report = constants.ValidationFail
		var reasons []string

		for _, fail := range fails {
			reasons = append(reasons, fmt.Sprintf("- %s", utils.Capitalize(fail.Error())))
		}
		reason = fmt.Sprintf(constants.ReasonTemplate, strings.Join(reasons, "\n"))
	}

	resultReport := fmt.Sprintf(constants.ResultTemplate, report)

	result := fmt.Sprintf(constants.ValidatorTemplate, strings.Join(resultList, "\n"), resultReport)

	if reason != "" {
		result = fmt.Sprintf("%s\n\n%s", result, reason)
	}

	return result
}

// FormatResult formats both whitelist and validation results for workflow reporting in markdown syntax
func FormatResult(
	whitelistResults []*entity.WhitelistResult,
	validationResults []*entity.ValidationResult,
) string {
	report := constants.ReportHeader

	report = fmt.Sprintf("%s\n\n%s", report, formatWhitelistResult(whitelistResults))

	if len(validationResults) > 0 {
		report = fmt.Sprintf("%s\n\n%s", report, formatValidationResult(validationResults))
	}

	return report
}
