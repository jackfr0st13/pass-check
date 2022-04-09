package pass

func LengthRule(minlength, maxlength int) CheckerFunction {
	return CheckerFunction(func(passData PasswordData) RuleResult {
		var rr RuleResult = *NewRuleResult()
		if len(passData.password) > maxlength {
			rr.AddError(string(ERR_CODE_MAX), createRuleResultParameterForLengthRule(minlength, maxlength))
		} else if len(passData.password) < minlength {
			rr.AddError(string(ERR_CODE_MIN), createRuleResultParameterForLengthRule(minlength, maxlength))
		}
		return rr
	})
}

func createRuleResultParameterForLengthRule(minlength, maxlength int) map[string]interface{} {
	var result map[string]interface{} = make(map[string]interface{})
	result["minimumLength"] = minlength
	result["maximumLength"] = maxlength
	return result
}
