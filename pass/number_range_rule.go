package pass

import (
	"strconv"
	"strings"
)

func NumberRangeRule(lowerRange, upperRange int, failAtFirstFailure bool) CheckerFunction {
	return CheckerFunction(func(passData PasswordData) RuleResult {
		var rr RuleResult = *NewRuleResult()
		for i := lowerRange; i <= upperRange; i++ {
			if strings.Contains(passData.password, strconv.Itoa(i)) {
				rr.AddError(string(ERR_CODE_ILLEGAL_NUMBER_RANGE), createRuleResultDetailParametersForNumberRangeRule(lowerRange, upperRange, i))
				if failAtFirstFailure {
					return rr
				}
			}
		}
		return rr
	})
}

func createRuleResultDetailParametersForNumberRangeRule(lowerRange, upperRange, number int) map[string]interface{} {
	var result map[string]interface{} = make(map[string]interface{})
	result["number"] = number
	result["matchBehavior"] = CONTAINS
	return result
}
