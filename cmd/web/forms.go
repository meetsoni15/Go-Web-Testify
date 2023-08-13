package main

import (
	"net/url"
	"strings"
)

// errors is convenience type, so we can have function tied to our map
type errors map[string][]string

// get error by passing field
func (e errors) Get(field string) string {
	errorSlice := e[field]
	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}

// add errors
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Form is the type used to instantiate the validation
type Form struct {
	Data   url.Values
	Errors errors
}

// NewForm is instantiate new validation
func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}

// Has check if form has field or not
func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	return x != ""
}

// Required fields are populated or not
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Check will check boolean condition and if failed
// So add error inside error map
func (f *Form) Check(ok bool, field, message string) {
	if !ok {
		f.Errors.Add(field, message)
	}
}

// Valid will return bool based on errors slice
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
