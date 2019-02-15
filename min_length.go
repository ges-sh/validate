package validate

import (
	"fmt"
	"unicode/utf8"
)

// MinLength is validator which requires string to contain at least Min characters.
type MinLength struct {
	Min int
	Err error
}

// Valid checks whether provided v is having at least Min characters. If v is not string, Valid will panic.
func (m MinLength) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("wrong data type, expected int, got: %T", v))
	}

	if utf8.RuneCountInString(val) < m.Min {
		return m.Err
	}

	return nil
}
