package validate

import (
	"fmt"
	"net/url"
)

// URL is validator which requires string to be valid url.URL. If scheme is not empty,
// URL scheme will be checked against it.
type URL struct {
	Scheme string
	Err    error
}

// Valid checks whether provided v is valid url. If v is not a string, Valid will panic.
func (r URL) Valid(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		panic(fmt.Sprintf("wrong data type, expected string, got: %T", v))
	}

	url, err := url.Parse(val)
	if err != nil {
		if r.Err != nil {
			return r.Err
		}
		return err
	}

	if r.Scheme != "" && url.Scheme != r.Scheme {
		return r.Err
	}

	return nil
}
