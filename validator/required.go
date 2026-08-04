package validator

import "reflect"

type Required struct {
	val any
}

func NewValidationRequired(val any) *Required {
	return &Required{val: val}
}
func (v *Required) Validate(path string, value any) bool {
	if v.val == nil {
		return false
	}

	val, wasNil := ResolveValue(v.val)
	if wasNil {
		return false
	}

	switch val.Kind() {
	case reflect.String:
		return len(val.String()) > 0
	case reflect.Slice, reflect.Array, reflect.Map:
		return val.Len() > 0
	case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return val.Int() != 0
	case reflect.Float32, reflect.Float64:
		return val.Float() != 0
	case reflect.Func:
		isNil := val.IsNil()
		return !isNil
	default:
		return true
	}
}
