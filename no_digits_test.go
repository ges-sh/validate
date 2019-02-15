package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestNoDigits(t *testing.T) {
	var errContainsDigits = errors.New("cannot contain digits")

	validator := NoDigits{Err: errContainsDigits}

	testCases := []struct {
		value  string
		expErr error
	}{
		{
			value:  "test",
			expErr: nil,
		},
		{
			value:  "test12",
			expErr: errContainsDigits,
		},
	}

	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Running case nr %d", i), func(t *testing.T) {
			err := validator.Valid(tt.value)
			if err != tt.expErr {
				t.Errorf("\nexp err: %v\ngot err: %v", tt.expErr, err)
				return
			}
		})
	}
}
