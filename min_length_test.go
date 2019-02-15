package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestMinLength(t *testing.T) {
	var errInvalidLength = errors.New("value is too short")

	validator := MinLength{Min: 3, Err: errInvalidLength}

	testCases := []struct {
		value  string
		expErr error
	}{
		{
			value:  "test",
			expErr: nil,
		},
		{
			value:  "te",
			expErr: errInvalidLength,
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
