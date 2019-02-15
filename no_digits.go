package validate

import (
	"fmt"
	"regexp"
)

var digitRegexp = regexp.MustCompile(`\d`)

// NoDigits is an validator which checks if provided value contains digits
type NoDigits struct {
	Err error
}

// Valid checks whether provided v contains digits. If v is not string, Valid will panic.
func (d NoDigits) Valid(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("expected a string, got %T", v))
	}

	if digitRegexp.MatchString(str) {
		return d.Err
	}

	return nil
}
