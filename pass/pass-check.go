package pass

import (
	"errors"
	"fmt"
)

type PassChecker []CheckerFunction

func NewPassChecker(vfunc ...CheckerFunction) *PassChecker {
	pc := PassChecker{}
	pc = append(pc, vfunc...)
	return &pc
}

func (pc *PassChecker) Check(passData PasswordData) (RuleResult, error) {
	if len(*pc) == 0 {
		return *NewRuleResult(WithValidValue(false)), errors.New("rules cannot be empty")
	}
	ruleResult := NewRuleResult()
	for _, rule := range *pc {
		rr := rule(passData)
		if !rr.IsValid() {
			ruleResult.SetValid(false)
			ruleResult.SetDetails(append(ruleResult.GetDetails(), rr.GetDetails()...))
		}
	}
	return *ruleResult, nil
}

func (pc *PassChecker) GetMessages(ruleResult RuleResult) []string {
	result := []string{}
	for _, detail := range ruleResult.GetDetails() {
		message := resolveMessage(detail)
		result = append(result, message)
	}
	return result
}

func resolveMessage(detail RuleResultDetail) string {
	message := ""
	for _, key := range detail.GetErrorCodes() {
		//fetch message from some map here
		message = errorValuesMap[key]
		if len(message) > 0 {
			return fmt.Sprintf(message, detail.GetValues()...)
		}
	}
	format := ""
	if len(detail.GetParameters()) > 0 {
		format = fmt.Sprint("%s : %v", detail.GetErrorCode(), detail.GetParameters())
	} else {
		format = detail.GetErrorCode()
	}
	return format
}
