package validator

import "reflect"

type Min struct {
	val any
	min int
}

func NewValidationMin(val any, vMin int) *Min {
	return &Min{val: val, min: vMin}
}

func (v *Min) Validate(path string, value any) bool {
	val, wasNil := ResolveValue(v.val)
	if wasNil {
		return false
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() >= int64(v.min)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() >= uint64(v.min)
	case reflect.Float32, reflect.Float64:
		return val.Float() >= float64(v.min)
	case reflect.String:
		return len(val.String()) >= v.min
	case reflect.Slice, reflect.Array, reflect.Map:
		return val.Len() >= v.min
	default:
		return false
	}
}
