package pass

import (
	"strings"

	"github.com/pkg/errors"
)

type RuleResultDetail struct {
	errorCodes []string
	parameters map[string]interface{}
}

func NewRuleResuleDetailWithCode(code string, params map[string]interface{}) (RuleResultDetail, error) {
	if len(strings.TrimSpace(code)) == 0 {
		return RuleResultDetail{}, errors.New("code must have a value")
	}
	if params == nil {
		return RuleResultDetail{
			errorCodes: []string{code},
			parameters: make(map[string]interface{}),
		}, nil
	}
	return RuleResultDetail{
		errorCodes: []string{code},
		parameters: params,
	}, nil
}

func NewRuleResuleDetailWithCodes(codes []string, params map[string]interface{}) (RuleResultDetail, error) {
	if len(codes) == 0 {
		return RuleResultDetail{}, errors.New("codes cannot be an empty array")
	}
	if params == nil {
		return RuleResultDetail{
			errorCodes: codes,
			parameters: make(map[string]interface{}),
		}, nil
	}
	return RuleResultDetail{
		errorCodes: codes,
		parameters: params,
	}, nil
}

func (rrd RuleResultDetail) GetErrorCodes() []string {
	return rrd.errorCodes
}

func (rrd RuleResultDetail) GetErrorCode() string {
	return rrd.errorCodes[len(rrd.errorCodes)-1]
}

func (rrd RuleResultDetail) GetParameters() map[string]interface{} {
	return rrd.parameters
}

func (rrd RuleResultDetail) GetValues() []interface{} {
	var arr []interface{}
	for _, v := range rrd.parameters {
		arr = append(arr, v)
	}
	return arr
}
