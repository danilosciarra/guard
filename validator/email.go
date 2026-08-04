package validator

import (
	"net/mail"
	"reflect"
)

type Email struct {
	val any
}

func NewValidationEmail(val any) *Email {
	return &Email{val: val}
}

func (v *Email) Validate(path string, value any) bool {
	val := reflect.ValueOf(v.val)

	if val.Kind() != reflect.String {
		return false
	}

	addr, err := mail.ParseAddress(val.String())
	return err == nil && addr.Name == ""
}
