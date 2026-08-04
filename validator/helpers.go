package validator

import "reflect"

func ResolveValue(v any) (reflect.Value, bool) {
	val := reflect.ValueOf(v)

	for val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return val, true
		}
		val = val.Elem()
	}

	if val.Kind() == reflect.Interface {
		if val.IsNil() {
			return val, true
		}
		val = val.Elem()
	}

	return val, false
}

