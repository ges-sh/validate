package validate

import (
	"fmt"
	"regexp"
)

// Email is validator which checks if provided value is valid email address. If no Regexp is specified, Email will use default one.
type Email struct {
	Regexp *regexp.Regexp
	Err    error
}

// EmailRegexp is used by Email if no other regexp is specified.
var EmailRegexp = regexp.MustCompile(
	// mailbox
	"^[a-z0-9!#$%&'*+/i=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*" +
		// @host
		"@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$",
)

// Valid checks whether provided string is valid email address. If v is not string, Valid will panic.
func (e Email) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("wrong data type, expected string, got: %T", v))
	}

	var valid bool
	if e.Regexp == nil {
		valid = EmailRegexp.MatchString(val)
	} else {
		valid = e.Regexp.MatchString(val)
	}

	if !valid {
		return e.Err
	}

	return nil
}
