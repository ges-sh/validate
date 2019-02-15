package validate

import (
	"fmt"
)

// Bool is validator which required boolean to be the same as Value
type Bool struct {
	Value bool
	Err   error
}

// Valid checks whether provided boolean is the same as the one specified by Value. If v is not boolean, Valid will panic.
func (b Bool) Valid(v interface{}) error {
	val, ok := v.(bool)
	if !ok {
		panic(fmt.Sprintf("wrong data type, expected bool, got: %T", v))
	}

	if val != b.Value {
		return b.Err
	}

	return nil
}
