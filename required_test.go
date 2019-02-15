package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestRequired(t *testing.T) {
	var errRequired = errors.New("value cannot be empty")

	validator := Required{Err: errRequired}

	testCases := []struct {
		value  string
		expErr error
	}{
		{
			value:  "required",
			expErr: nil,
		},
		{
			value:  "",
			expErr: errRequired,
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
