package validator

import (
	"fmt"
	"regexp"
)

type Regex struct {
	val   any
	regex string
}

func NewValidationRegex(val any, vRegex string) *Regex {
	return &Regex{val: val, regex: vRegex}
}
func (v *Regex) Validate(path string, value any) bool {
	vString := fmt.Sprintf("%v", v.val)
	matched, err := regexp.Match(v.regex, []byte(vString))
	if err != nil {
		return false
	}
	return matched
}
