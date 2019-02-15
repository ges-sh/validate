package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestNoLetters(t *testing.T) {
	var errContainsLetters = errors.New("cannot contain letters")

	validator := NoLetters{Err: errContainsLetters}

	testCases := []struct {
		value  string
		expErr error
	}{
		{
			value:  "12345",
			expErr: nil,
		},
		{
			value:  "12345as",
			expErr: errContainsLetters,
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
