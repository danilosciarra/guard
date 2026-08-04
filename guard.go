package guard

import (
	"fmt"
	"github.com/danilosciarra/guard/validator"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Guard struct {
	validatorTags            map[string]validationTag
	obj                      any
	customValidatorFunctions map[string]customValidationFunction
	isSlice                  bool
}

type customValidationFunction func(path string, value any) bool

type validationTag struct {
	effectivePath string
	tag           string
}

func New() *Guard {
	return &Guard{}
}
func (m *Guard) RegisterValidator(tag string, val customValidationFunction) {
	if m.customValidatorFunctions == nil {
		m.customValidatorFunctions = make(map[string]customValidationFunction)
	}
	m.customValidatorFunctions[tag] = val
}

func (m *Guard) resolveValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return reflect.Value{}
		}
		val = val.Elem()
	}

	if val.Kind() == reflect.Interface {
		if !val.IsNil() {
			val = reflect.ValueOf(val.Interface())
			if val.Kind() == reflect.Pointer {
				val = val.Elem()
			}
		}
	}
	return val
}

func (m *Guard) Validate(obj any) error {
	if obj == nil {
		return &InvalidInputError{Reason: "object cannot be nil"}
	}

	g := &Guard{
		obj:                      obj,
		customValidatorFunctions: m.customValidatorFunctions,
	}

	if err := g.buildValidationTags(obj, ""); err != nil {
		return err
	}

	if g.isSlice {
		myObj := g.obj
		if reflect.TypeOf(myObj).Kind() == reflect.Pointer {
			myObj = reflect.ValueOf(g.obj).Elem().Interface()
		}
		sObj := reflect.ValueOf(myObj)
		for i := 0; i < sObj.Len(); i++ {
			item := sObj.Index(i).Interface()
			itemManager := New()
			itemManager.customValidatorFunctions = m.customValidatorFunctions
			if err := itemManager.Validate(item); err != nil {
				return errors.Wrapf(err, "item %d validation failed", i)
			}
		}
		return nil
	}

	for key, tag := range g.validatorTags {
		validationParts := strings.Split(tag.tag, ",")
		val, ok := g.getValueByPath(tag.effectivePath)
		if !ok {
			return &PathError{
				Path:   tag.effectivePath,
				Reason: fmt.Sprintf("field not found in object (key: %s)", key),
			}
		}
		for _, validationTag := range validationParts {
			validationTag = strings.TrimSpace(validationTag)
			valid := true
			if customValidationFunc, ok := g.customValidatorFunctions[validationTag]; ok {
				valid = customValidationFunc(tag.effectivePath, val)
			} else {
				v, err := validator.New(validationTag, val)
				if err != nil {
					return errors.Wrapf(err, "field '%s'", tag.effectivePath)
				}
				if v != nil {
					valid = v.Validate(tag.effectivePath, val)
				}
			}
			if !valid {
				displayVal := val
				if val == "" {
					displayVal = "(empty)"
				}
				return &ValidationError{
					FieldPath: tag.effectivePath,
					Value:     displayVal,
					Tag:       validationTag,
					Reason:    "validation constraint not met",
				}
			}
		}

	}
	return nil
}

func (m *Guard) buildValidationTags(obj interface{}, basePath string) error {
	if obj == nil {
		return &InvalidInputError{Reason: "cannot build validation tags for nil object"}
	}

	val := reflect.ValueOf(obj)
	val = m.resolveValue(val)
	if !val.IsValid() {
		return &InvalidInputError{Reason: fmt.Sprintf("cannot resolve value at path '%s'", basePath)}
	}

	if val.Kind() == reflect.Slice {
		if val.Len() == 0 {
			return nil
		}
		val = val.Index(0)
		if val.Kind() == reflect.Pointer {
			if val.IsNil() {
				return nil
			}
			val = val.Elem()
		}
		m.isSlice = true
	}
	if val.Kind() != reflect.Struct {
		return &TypeMismatchError{
			Expected: "struct or slice",
			Actual:   val.Kind().String(),
			Field:    basePath,
		}
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)
		pathKey := field.Name
		if basePath != "" {
			pathKey = fmt.Sprintf("%s/%s", basePath, field.Name)
		}
		if fieldVal.Kind() == reflect.Pointer {
			if fieldVal.IsNil() {
				continue
			}
			fieldVal = fieldVal.Elem()
		}
		if fieldVal.Kind() == reflect.Struct {
			if err := m.buildValidationTags(fieldVal.Interface(), pathKey); err != nil {
				return err
			}
		}
		if fieldVal.Kind() == reflect.Slice {
			for j := 0; j < fieldVal.Len(); j++ {
				item := fieldVal.Index(j)
				if item.Kind() == reflect.Pointer {
					if item.IsNil() {
						continue
					}
					item = item.Elem()
				}
				if item.Kind() != reflect.Struct {
					continue
				}
				iName := strconv.Itoa(j)
				idVal := item.FieldByName("Id")
				if idVal.IsValid() && idVal.Kind() == reflect.String && idVal.String() != "" {
					iName = idVal.String()
				}
				if err := m.buildValidationTags(fieldVal.Index(j).Interface(), fmt.Sprintf("%s/%s", pathKey, iName)); err != nil {
					return err
				}
			}
		}
		tag := field.Tag.Get("validate")
		if tag != "" {
			if m.validatorTags == nil {
				m.validatorTags = make(map[string]validationTag)
			}
			m.validatorTags[strings.ToLower(pathKey)] = validationTag{
				effectivePath: pathKey,
				tag:           tag,
			}
		}
	}
	return nil
}
func (m *Guard) getValueByPath(path string) (interface{}, bool) {
	parts := strings.Split(path, "/")
	val := reflect.ValueOf(m.obj)
	val = m.resolveValue(val)
	if !val.IsValid() {
		return nil, false
	}

	for _, part := range parts {
		if val.Kind() == reflect.Struct {
			val = val.FieldByName(part)
			if !val.IsValid() {
				return nil, false
			}
		} else if val.Kind() == reflect.Slice {
			index := -1
			for i := 0; i < val.Len(); i++ {
				item := val.Index(i)
				for item.Kind() == reflect.Pointer {
					if item.IsNil() {
						break
					}
					item = item.Elem()
				}
				if item.Kind() != reflect.Struct {
					continue
				}
				idVal := item.FieldByName("Id")
				if idVal.IsValid() && idVal.Kind() == reflect.String && idVal.String() == part {
					index = i
					break
				}
			}
			if index == -1 {
				intIndex, err := strconv.Atoi(part)
				if err != nil || intIndex < 0 || intIndex >= val.Len() {
					return nil, false
				}
				index = intIndex
			}
			val = val.Index(index)
		} else {
			return nil, false
		}
		for val.Kind() == reflect.Pointer {
			if val.IsNil() {
				return nil, false
			}
			val = val.Elem()
		}

	}
	return val.Interface(), true
}
