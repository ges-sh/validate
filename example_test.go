package validate_test

import (
	"errors"
	"fmt"

	"github.com/ges-sh/validate"
)

type str struct {
	FirstName string `valid:"firstName"`
	LastName  string `valid:"lastName"`
}

var validators = validate.Validators{
	"firstName": []validate.Validator{
		validate.Required{Err: errors.New("cannot be empty")},
	},
	"lastName": []validate.Validator{
		validate.Required{Err: errors.New("cannot be empty")},
	},
}

func Example() {
	data := str{
		FirstName: "John",
		LastName:  "Smith",
	}

	errors := validators.Validate(data)
	fmt.Println(errors) // map[]
}
