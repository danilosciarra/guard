# Guard

[![Go Reference](https://pkg.go.dev/badge/github.com/danilosciarra/guard.svg)](https://pkg.go.dev/github.com/danilosciarra/guard)
[![License: MIT](https://img.shields.io/github/license/danilosciarra/guard)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/danilosciarra/guard)](go.mod)

A lightweight, extensible validation library for Go with first-class support for nested structs, slices, and custom validators.

## Features

- ✅ Zero dependencies
- ✅ Built-in validators (`required`, `min`, `max`, `regex`, `oneOf`, `email`)
- ✅ Custom validators
- ✅ Nested struct validation
- ✅ Slice validation
- ✅ Typed validation errors

## Installation

```bash
go get github.com/danilosciarra/guard
```

```go
import "github.com/danilosciarra/guard"
```

## Usage

```go
v := guard.New()

v.RegisterValidator(...)

err := v.Validate(&obj)
```

### Basic usage

```go
package main

import (
	"fmt"

	"github.com/danilosciarra/guard"
)

type User struct {
	Name  string `validate:"required,min=3"`
	Age   int    `validate:"min=18,max=120"`
	Email string `validate:"required,email"`
}

func main() {
	u := User{Name: "Dan", Age: 30, Email: "dan@example.com"}

	v := guard.New()
	ok, err := v.Validate(&u)
	if err != nil {
		fmt.Println("validation error:", err)
		return
	}

	fmt.Println("valid:", ok)
}
```

## Validation tags

Use the `validate` struct tag. Multiple rules are comma-separated.

```go
Field string `validate:"required,min=3,max=20"`
```

### Supported tags

| Tag | Syntax | Applies to | Description |
|-----|--------|------------|-------------|
| `required` | `required` | `string`, `int`, `float`, `slice`, `array`, `map`, `func` | Fails if zero/empty/nil |
| `min` | `min=<int>` | `int`, `float`, `string`, `slice`, `array`, `map` | Value/length must be >= min |
| `max` | `max=<int>` | `int`, `float`, `string`, `slice`, `array`, `map` | Value/length must be <= max |
| `regex` | `regex=<pattern>` | any (converted to string) | Must match the Go regexp pattern |
| `oneOf` | `oneOf=a\|b\|c` | `string`, `int`, `float` | Value must be one of the pipe-separated options |
| `email` | `email` | `string` | Must be a valid email address (no display name) |

#### Type support matrix

| Tag | string | int/uint | float | slice/array/map | func | pointer |
|-----|:------:|:--------:|:-----:|:---------------:|:----:|:-------:|
| `required` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ (dereferenced) |
| `min` | ✓ (len) | ✓ | ✓ | ✓ (len) | — | ✓ (dereferenced) |
| `max` | ✓ (len) | ✓ | ✓ | ✓ (len) | — | ✓ (dereferenced) |
| `regex` | ✓ | ✓ (fmt) | ✓ (fmt) | — | — | — |
| `oneOf` | ✓ | ✓ (fmt) | ✓ (fmt) | — | — | ✓ (dereferenced) |
| `email` | ✓ | — | — | — | — | — |

> Pointers and interfaces are automatically dereferenced before validation. Nil pointers return `false` for `required` and all numeric/string validators.

## Multiple rules on the same field

```go
type Product struct {
	Code  string  `validate:"required,regex=^[A-Z0-9_-]+$"`
	Price float64 `validate:"min=1,max=9999"`
}
```

Rules are evaluated left-to-right. Validation stops at the first failing rule.

## Custom validators

You can register custom validators by tag name.

```go
type Payload struct {
	Token string `validate:"token_prefix"`
}

v := guard.New()
v.RegisterValidator("token_prefix", func(path string, value any) bool {
	s, ok := value.(string)
	if !ok {
		return false
	}
	return len(s) > 4 && s[:4] == "tok_"
})

ok, err := v.Validate(&Payload{Token: "tok_123"})
```

If a custom validator is registered for a tag, it is used instead of built-in validator parsing.

## Nested structs and slices

The package traverses:

- nested structs
- pointers to structs (nil pointers are skipped unless a direct tagged field is evaluated)
- slices of structs

For slices of structs, when items have an `Id string` field, that value may be used internally for path resolution.

### Example with nested data

```go
type Address struct {
	City string `validate:"required"`
}

type Customer struct {
	Id      string   `validate:"required"`
	Name    string   `validate:"required,min=2"`
	Address *Address `validate:"required"`
}

type Order struct {
	Items []Customer
}

order := Order{
	Items: []Customer{
		{Id: "c1", Name: "Alice", Address: &Address{City: "Rome"}},
	},
}

v := guard.New()
ok, err := v.Validate(&order)
```

## Error handling

`Validate` returns `(bool, error)`.

- `ok == true, err == nil`: all validations passed
- `ok == false, err != nil`: validation failed or invalid configuration/input

The package defines typed errors in `errors.go`:

- `*guard.ValidationError`
- `*guard.PathError`
- `*guard.TypeMismatchError`
- `*guard.InvalidInputError`

Validator parsing errors come from the `validator` package and are wrapped with context.

### Example: inspect error type

```go
ok, err := v.Validate(&u)
if err != nil {
	switch e := err.(type) {
	case *guard.ValidationError:
		fmt.Println("field:", e.FieldPath, "tag:", e.Tag, "value:", e.Value)
	case *guard.InvalidInputError:
		fmt.Println("invalid input:", e.Reason)
	default:
		fmt.Println("error:", err)
	}
}
```

## Common pitfalls

- Unknown tag names return validation/parsing errors.
- `min`/`max` parameters must be integers.
- `email` validator expects a string value.
- `regex` uses Go regexp syntax.


## License

This project is distributed under the terms in `LICENSE`.

