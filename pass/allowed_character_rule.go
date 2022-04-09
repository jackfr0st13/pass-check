package pass

func AllowedCharactersRule(allowerdChars []rune, matchBehaviour MatchBehaviour) CheckerFunction {
	return CheckerFunction(func(passData PasswordData) RuleResult {
		var rr RuleResult = *NewRuleResult()
		return rr
	})
}
