package validate

import (
	"errors"
	"testing"
)

type str struct {
	FirstName string `valid:"firstName"`
	LastName  string `valid:"lastName"`
}

var validators = Validators{
	"firstName": []Validator{
		Required{Err: errors.New("cannot be empty")},
		MinLength{Min: 2, Err: errors.New("too short")},
	},
	"lastName": []Validator{
		Required{Err: errors.New("cannot be empty")},
	},
}

func TestValidate(t *testing.T) {
	exp := Errors{
		"firstName": []error{errors.New("cannot be empty"), errors.New("too short")},
		"lastName":  []error{errors.New("cannot be empty")},
	}

	// missing first name
	data := str{}

	errors := validators.Validate(data)

	for field, errs := range errors {
		for i, err := range errs {
			if expErr := exp[field][i].Error(); expErr != err.Error() {
				t.Errorf("\nExp: %v\nGot: %v\n", expErr, err)
			}
		}
	}
}
