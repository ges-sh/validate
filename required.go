package validate

import (
	"fmt"
)

// Required is validator which requires string to not be empty
type Required struct {
	Err error
}

// Valid checks whether provided v is empty. If v is not a string, Valid will panic.
func (r Required) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("wrong data type, expected string, got: %T", v))
	}

	if val == "" {
		return r.Err
	}

	return nil
}
