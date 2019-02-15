package validate

import (
	"fmt"
	"regexp"
)

// Phone is validator which checks is provided value is valid phone number. If no Regexp is specified, Phone will use default E.164 regexp.
type Phone struct {
	Regexp *regexp.Regexp
	Err    error
}

// E164Regexp represents E.164 phone number regexp.
var E164Regexp = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

// Valid checks whether provided string is valid email address. If v is not string, Valid will panic.
func (p Phone) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("wront data type, expected int, got: %T", v))
	}

	var valid bool
	if p.Regexp == nil {
		valid = E164Regexp.MatchString(val)
	} else {
		valid = p.Regexp.MatchString(val)
	}

	if !valid {
		return p.Err
	}

	return nil
}
