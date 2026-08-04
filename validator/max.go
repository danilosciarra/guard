package validator

import "reflect"

type Max struct {
	val any
	max int
}

func NewValidationMax(val any, vMax int) *Max {
	return &Max{val: val, max: vMax}
}
func (v *Max) Validate(path string, value any) bool {
	val, wasNil := ResolveValue(v.val)
	if wasNil {
		return false
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() <= int64(v.max)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() <= uint64(v.max)
	case reflect.Float32, reflect.Float64:
		return val.Float() <= float64(v.max)
	case reflect.String:
		return len(val.String()) <= v.max
	case reflect.Slice, reflect.Array, reflect.Map:
		return val.Len() <= v.max
	default:
		return false
	}
}
