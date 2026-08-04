package validator

import (
	"reflect"
	"strconv"
)

type OneOf struct {
	val any
	of  []string
}

func NewValidationOneOf(val any, of []string) *OneOf {
	return &OneOf{val: val, of: of}
}

func (v *OneOf) Validate(path string, value any) bool {
	val, wasNil := ResolveValue(v.val)
	if wasNil {
		return false
	}
	strToChek := ""
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		strToChek = strconv.Itoa(int(val.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		strToChek = strconv.Itoa(int(val.Uint()))
	case reflect.Float64:
		strToChek = strconv.FormatFloat(val.Float(), 'f', -1, 64)
	case reflect.Float32:
		strToChek = strconv.FormatFloat(val.Float(), 'f', -1, 32)
	case reflect.String:
		strToChek = val.String()
	default:
		return false
	}
	if len(v.of) == 0 {
		return true
	}
	for _, allowedValue := range v.of {
		if strToChek == allowedValue {
			return true
		}
	}
	return false
}
