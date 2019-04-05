package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestURL(t *testing.T) {
	var errInvalidURL = errors.New("invalid url")

	testCases := []struct {
		url       string
		expErr    error
		validator URL
	}{
		{
			url:       "https://example.com",
			expErr:    nil,
			validator: URL{Err: errInvalidURL},
		},
		{
			url:       "$@$!%@!",
			expErr:    errInvalidURL,
			validator: URL{Err: errInvalidURL},
		},
		{
			url:       "ftp://example.com",
			expErr:    errInvalidURL,
			validator: URL{Err: errInvalidURL, Scheme: "https"},
		},
	}

	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Running case nr %d", i), func(t *testing.T) {
			err := tt.validator.Valid(tt.url)
			if err != tt.expErr {
				t.Errorf("\nexp err: %v\ngot err: %v", tt.expErr, err)
				return
			}
		})
	}
}
