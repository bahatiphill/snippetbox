package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

//This function initialize a custum Form struct
func New(data url.Values) *Form {
	return &Form {
		data,
		errors(map[string][]string{}),
	}
}

// Function to check if field is not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be Empty")
		}
	}
}

//check the lenght of the provided data
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long(maximum is %d characters)", d))
	}
}

//check permitted values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This Field is Invalid")
}

// a valid method which return true
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}