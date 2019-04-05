package validate

import (
	"fmt"
	"unicode/utf8"
)

// MaxLength is validator which required string to contain at most Max characters.
type MaxLength struct {
	Max int
	Err error
}

// Valid checks whether provided v is having at most Max characters. If v is not string, Valid will panic.
func (l MaxLength) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("wrong data type, expected string, got: %T", v))
	}

	if utf8.RuneCountInString(val) > l.Max {
		return l.Err
	}

	return nil
}
