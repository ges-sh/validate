package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	var errExpectedTrue = errors.New("expected bool to be true")

	validator := Bool{Value: true, Err: errExpectedTrue}

	testCases := []struct {
		agreed bool
		expErr error
	}{
		{
			agreed: true,
			expErr: nil,
		},
		{
			agreed: false,
			expErr: errExpectedTrue,
		},
	}

	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Running case nr %d", i), func(t *testing.T) {
			err := validator.Valid(tt.agreed)
			if err != tt.expErr {
				t.Errorf("\nexp err: %v\ngot err: %v", tt.expErr, err)
				return
			}
		})
	}
}
