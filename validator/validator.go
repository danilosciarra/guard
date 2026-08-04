package validator

import (
	"fmt"
	"strconv"
	"strings"
)

type Type string

const (
	TypeMin      Type = "min"
	TypeMax      Type = "max"
	TypeRequired Type = "required"
	TypeRegex    Type = "regex"
	TypeOneOf    Type = "oneOf"
	TypeEmail    Type = "email"
)

type validator interface {
	Validate(path string, value any) bool
}

type ParseError struct {
	Tag    string
	Reason string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("tag parse error: tag '%s' - %s", e.Tag, e.Reason)
}

func New(tag string, val any) (validator, error) {
	if tag == "" {
		return nil, &ParseError{Tag: tag, Reason: "tag cannot be empty"}
	}

	vType, param, err := parseTagParameter(tag)
	if err != nil {
		return nil, err
	}

	switch vType {
	case TypeMin:
		if param == nil {
			return nil, &ParseError{Tag: tag, Reason: "min tag requires a numeric parameter"}
		}
		return NewValidationMin(val, param.(int)), nil
	case TypeMax:
		if param == nil {
			return nil, &ParseError{Tag: tag, Reason: "max tag requires a numeric parameter"}
		}
		return NewValidationMax(val, param.(int)), nil
	case TypeRequired:
		return NewValidationRequired(val), nil
	case TypeRegex:
		if param == nil {
			return nil, &ParseError{Tag: tag, Reason: "regex tag requires a pattern parameter"}
		}
		return NewValidationRegex(val, param.(string)), nil
	case TypeOneOf:
		if param == nil {
			return nil, &ParseError{Tag: tag, Reason: "oneOf tag requires a pipe-separated values parameter"}
		}
		return NewValidationOneOf(val, strings.Split(param.(string), "|")), nil
	case TypeEmail:
		return NewValidationEmail(val), nil
	default:
		return nil, &ParseError{Tag: tag, Reason: fmt.Sprintf("unknown tag type '%s'", vType)}
	}
}

func parseTagParameter(tag string) (Type, any, error) {
	var vType Type
	parts := strings.Split(tag, "=")
	
	if len(parts) == 0 || parts[0] == "" {
		return "", nil, &ParseError{Tag: tag, Reason: "tag type is missing"}
	}

	vType = Type(parts[0])

	if !isValidType(vType) {
		return "", nil, &ParseError{Tag: tag, Reason: fmt.Sprintf("unknown tag type '%s'", vType)}
	}

	var param any
	var err error

	switch vType {
	case TypeMin, TypeMax:
		if len(parts) < 2 || parts[1] == "" {
			return vType, nil, &ParseError{Tag: tag, Reason: fmt.Sprintf("%s tag requires a numeric parameter", vType)}
		}
		param, err = strconv.Atoi(parts[1])
		if err != nil {
			return vType, nil, &ParseError{Tag: tag, Reason: fmt.Sprintf("%s parameter must be a valid integer, got '%s'", vType, parts[1])}
		}
	case TypeRegex:
		if len(parts) < 2 || parts[1] == "" {
			return vType, nil, &ParseError{Tag: tag, Reason: "regex tag requires a pattern parameter"}
		}
		param = parts[1]
	case TypeRequired:
		param = nil
	case TypeOneOf:
		if len(parts) < 2 || parts[1] == "" {
			return vType, nil, &ParseError{Tag: tag, Reason: "oneOf tag requires pipe-separated values"}
		}
		param = parts[1]
	case TypeEmail:
		param = nil
	}

	return vType, param, nil
}

func isValidType(t Type) bool {
	switch t {
	case TypeMin, TypeMax, TypeRequired, TypeRegex, TypeOneOf, TypeEmail:
		return true
	default:
		return false
	}
}
