package validate

import (
	"fmt"
	"reflect"
	"regexp"
)

var lettersRegex = regexp.MustCompile(`[a-zA-Z]`)

// NoLetters is an validator which checks if provided value contains letters
type NoLetters struct {
	Err error
}

// Valid checks whether provided v contains letters. If v is not string, Valid will panic.
func (l NoLetters) Valid(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		panic(fmt.Sprintf("expected a string, got %s", reflect.TypeOf(value)))
	}

	if lettersRegex.MatchString(str) {
		return l.Err
	}

	return nil
}
