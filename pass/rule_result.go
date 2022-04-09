package pass

import "github.com/pkg/errors"

type RuleResultOption func(*RuleResult)

type RuleResult struct {
	valid   bool
	details []RuleResultDetail
}

func NewRuleResult(opts ...RuleResultOption) *RuleResult {
	ruleResult := RuleResult{
		valid: true,
	}
	for _, o := range opts {
		o(&ruleResult)
	}
	return &ruleResult
}

func WithValidValue(value bool) RuleResultOption {
	return func(rr *RuleResult) {
		rr.valid = value
	}
}

func WithRuleResultDetails(details []RuleResultDetail) RuleResultOption {
	return func(rr *RuleResult) {
		rr.details = details
	}
}

func (rr *RuleResult) IsValid() bool {
	return rr.valid
}

func (rr *RuleResult) SetValid(value bool) {
	rr.valid = value
}

func (rr *RuleResult) GetDetails() []RuleResultDetail {
	return rr.details
}

func (rr *RuleResult) SetDetails(details []RuleResultDetail) {
	rr.details = details
}

func (rr *RuleResult) AddError(code string, params map[string]interface{}) error {
	rrd, err := NewRuleResuleDetailWithCode(code, params)
	if err != nil {
		return errors.Wrap(err, "error not added due to invalid code value")
	}
	rr.valid = false
	rr.details = append(rr.details, rrd)
	return nil
}

func (rr *RuleResult) AddErrors(codes []string, params map[string]interface{}) error {
	rrd, err := NewRuleResuleDetailWithCodes(codes, params)
	if err != nil {
		return errors.Wrap(err, "errors not added due to invalid codes value")
	}
	rr.valid = false
	rr.details = append(rr.details, rrd)
	return nil
}
