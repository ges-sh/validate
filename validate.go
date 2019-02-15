// Package validate allows user to validate Go structs using map of validators.
// Validating produces nice map of errors which can then be passed to for example frontend in the JSON format.
package validate

import (
	"reflect"
)

// Validator is checking whether v is passing all restrictions provided by Validator
type Validator interface {
	Valid(v interface{}) error
}

// Validators is a map of fields with theirs validators.
type Validators map[string][]Validator

// Errors represents list of errors within validated object
type Errors map[string][]error

// Validate checks whether v is valid by checking it with every provided Validator.
// If provided v is not a struct, it'll return empty errors map. If v contains
// not exported fields, Validate will panic, so it's caller responsibility to make sure
// all v fields are exported.
func (v Validators) Validate(s interface{}) Errors {
	errors := make(Errors)
	v.iterateFields(s, errors)
	return errors
}

func (v Validators) iterateFields(s interface{}, errors Errors) {
	valueOf := reflect.ValueOf(s)
	if valueOf.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < valueOf.NumField(); i++ {
		tag := valueOf.Type().Field(i).Tag.Get("valid")
		val := valueOf.Field(i)

		switch val.Kind() {
		case reflect.Struct:
			v.iterateFields(val.Interface(), errors)
			continue
		case reflect.Ptr:
			val = val.Elem()
		}

		v.validateField(tag, val, errors)
	}
}

func (v Validators) validateField(tag string, val reflect.Value, errors Errors) {
	for _, validator := range v[tag] {
		if err := validator.Valid(val.Interface()); err != nil {
			errors[tag] = append(errors[tag], err)
		}
	}
}
