package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	var errInvalidEmail = errors.New("invalid email address")

	validator := Email{Err: errInvalidEmail}

	testCases := []struct {
		email  string
		expErr error
	}{
		{
			email:  "user@test.com",
			expErr: nil,
		},
		{
			email:  "123.213@.com",
			expErr: errInvalidEmail,
		},
	}

	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Running case nr %d", i), func(t *testing.T) {
			err := validator.Valid(tt.email)
			if err != tt.expErr {
				t.Errorf("\nexp err: %v\ngot err: %v", tt.expErr, err)
				return
			}
		})
	}
}
