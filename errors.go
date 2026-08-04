package guard

import "fmt"

type ValidationError struct {
	FieldPath string
	Value     any
	Tag       string
	Reason    string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: field '%s' has value '%v' (tag: %s) - %s", e.FieldPath, e.Value, e.Tag, e.Reason)
}

type PathError struct {
	Path   string
	Reason string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("path error: path '%s' is invalid - %s", e.Path, e.Reason)
}

type TypeMismatchError struct {
	Expected string
	Actual   string
	Field    string
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch: field '%s' expected type %s but got %s", e.Field, e.Expected, e.Actual)
}

type InvalidInputError struct {
	Reason string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Reason)
}

type TagParseError struct {
	Tag    string
	Reason string
}

func (e *TagParseError) Error() string {
	return fmt.Sprintf("tag parsing error: tag '%s' is invalid - %s", e.Tag, e.Reason)
}

