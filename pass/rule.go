package pass

type Rule interface {
	Validate(passData PasswordData) RuleResult
}

type Validator interface {
	Rule
}
