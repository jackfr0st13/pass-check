package pass

import "unicode"

func WhitespaceRule(stopAtFailure bool) CheckerFunction {
	return CheckerFunction(func(passData PasswordData) RuleResult {
		var rr RuleResult = *NewRuleResult()
		for _, v := range passData.password {
			if unicode.IsSpace(v) {
				rr.AddError(string(ERR_CODE_ILLEGAL_WHITESPACE), createRuleResultParameterForWhitespaceRule(stopAtFailure, v))
				if stopAtFailure {
					return rr
				}
			}
		}
		return rr
	})
}

func createRuleResultParameterForWhitespaceRule(stopAtFailure bool, whitespaceCharacter rune) map[string]interface{} {
	var result map[string]interface{} = make(map[string]interface{})
	result["whitespaceCharacter"] = whitespaceCharacter
	result["matbehavior"] = CONTAINS
	return result
}
