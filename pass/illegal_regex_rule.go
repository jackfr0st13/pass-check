package pass

import "regexp"

func IllegalRegexRule(pattern string) CheckerFunction {
	return CheckerFunction(func(passData PasswordData) RuleResult {
		var rr RuleResult = *NewRuleResult()
		regex, err := regexp.Compile(pattern)
		if err != nil {
			rr.valid = false
			rr.AddError(string(ERR_CODE_ALLOWED_REGEX), createRuleResultDetailParametersForIllegalRegexRule(pattern))
		} else if matched := regex.MatchString(passData.password); !matched {
			rr.AddError(string(ERR_CODE_ALLOWED_REGEX), createRuleResultDetailParametersForIllegalRegexRule(pattern))
		}
		return rr
	})
}

func createRuleResultDetailParametersForIllegalRegexRule(pattern string) map[string]interface{} {
	var result map[string]interface{} = make(map[string]interface{})
	result["pattern"] = pattern
	return result
}
