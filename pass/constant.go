package pass

type ERR_CODE string

const (
	ERR_CODE_MIN                       ERR_CODE = "TOO_SHORT"
	ERR_CODE_MAX                       ERR_CODE = "TOO_LONG"
	ERR_CODE_ALLOWED_REGEX             ERR_CODE = "ALLOWED_MATCH"
	ERR_CODE_ILLEGAL_USERNAME          ERR_CODE = "ILLEGAL_USERNAME"
	ERR_CODE_ILLEGAL_USERNAME_REVERSED ERR_CODE = "ILLEGAL_USERNAME_REVERSED"
	ERR_CODE_ILLEGAL_WHITESPACE        ERR_CODE = "ILLEGAL_WHITESPACE"
	ERR_CODE_ILLEGAL_MATCH             ERR_CODE = "ILLEGAL_MATCH"
	ERR_CODE_ILLEGAL_NUMBER_RANGE      ERR_CODE = "ILLEGAL_NUMBER_RANGE ERR_CODE"
)

var errorValuesMap map[string]string = map[string]string{
	string(ERR_CODE_MIN):                       "Password must be %[1]d or more characters in length.",
	string(ERR_CODE_MAX):                       "Password must be no more than %[2]d characters in length.",
	string(ERR_CODE_ALLOWED_REGEX):             "Password must match pattern %[1]s",
	string(ERR_CODE_ILLEGAL_USERNAME):          "Password %[2]s the user id %[1]s",
	string(ERR_CODE_ILLEGAL_USERNAME_REVERSED): "Password %[2]s the user id %[1]s in reverse",
	string(ERR_CODE_ILLEGAL_WHITESPACE):        "Password %[2]s a whitespace character.",
	string(ERR_CODE_ILLEGAL_MATCH):             "Password matches the illegal pattern '%[1]s'",
	string(ERR_CODE_ILLEGAL_NUMBER_RANGE):      "Password %[2]s the number '%[1]s'.",
}
