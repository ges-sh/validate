package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestPhone(t *testing.T) {
	var errInvalidPhoneNumber = errors.New("invalid phone number")

	validator := Phone{Err: errInvalidPhoneNumber}

	testCases := []struct {
		number string
		expErr error
	}{
		{
			number: "+48666666666",
			expErr: nil,
		},
		{
			number: "invalid",
			expErr: errInvalidPhoneNumber,
		},
	}

	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Running case nr %d", i), func(t *testing.T) {
			err := validator.Valid(tt.number)
			if err != tt.expErr {
				t.Errorf("\nexp err: %v\ngot err: %v", tt.expErr, err)
				return
			}
		})
	}
}
