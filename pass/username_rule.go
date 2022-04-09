package pass

import (
	"strings"
)

func UsernameRule(ignoreCase, matchBackwards bool) CheckerFunction {
	return CheckerFunction(func(passData PasswordData) RuleResult {
		var rr RuleResult = *NewRuleResult()
		if len(passData.username) > 0 {
			if ignoreCase {
				if strings.Contains(strings.ToLower(passData.password), strings.ToLower(passData.username)) {
					//failure case
					rr.AddError(string(ERR_CODE_ILLEGAL_USERNAME), createRuleResultDetailParameterForUsernameRule(passData.username))
				}
			}
			if matchBackwards {
				reverseUsername := reverseString(passData.username)
				if strings.Contains(strings.ToLower(passData.password), strings.ToLower(reverseUsername)) {
					//failure case
					rr.AddError(string(ERR_CODE_ILLEGAL_USERNAME_REVERSED), createRuleResultDetailParameterForUsernameRule(passData.username))
				}
			}
		}
		return rr
	})
}

func createRuleResultDetailParameterForUsernameRule(username string) map[string]interface{} {
	var result map[string]interface{} = make(map[string]interface{})
	result["username"] = username
	result["matchbehaviour"] = CONTAINS
	return result
}
