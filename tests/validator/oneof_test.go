package validator

import (
	"github.com/danilosciarra/guard/validator"
	"testing"
)

func TestOneOfValidateInt(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "int valid value",
			val:      42,
			of:       []string{"42", "100", "200"},
			expected: true,
		},
		{
			name:     "int invalid value",
			val:      99,
			of:       []string{"42", "100", "200"},
			expected: false,
		},
		{
			name:     "int8 valid",
			val:      int8(10),
			of:       []string{"10", "20"},
			expected: true,
		},
		{
			name:     "int16 valid",
			val:      int16(1000),
			of:       []string{"1000", "2000"},
			expected: true,
		},
		{
			name:     "int32 valid",
			val:      int32(50000),
			of:       []string{"50000"},
			expected: true,
		},
		{
			name:     "int64 valid",
			val:      int64(999999),
			of:       []string{"999999"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOneOfValidateUint(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "uint valid value",
			val:      uint(42),
			of:       []string{"42", "100"},
			expected: true,
		},
		{
			name:     "uint invalid value",
			val:      uint(99),
			of:       []string{"42", "100"},
			expected: false,
		},
		{
			name:     "uint8 valid",
			val:      uint8(10),
			of:       []string{"10"},
			expected: true,
		},
		{
			name:     "uint16 valid",
			val:      uint16(1000),
			of:       []string{"1000"},
			expected: true,
		},
		{
			name:     "uint32 valid",
			val:      uint32(50000),
			of:       []string{"50000"},
			expected: true,
		},
		{
			name:     "uint64 valid",
			val:      uint64(999999),
			of:       []string{"999999"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOneOfValidateFloat(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "float32 valid value",
			val:      float32(3.14),
			of:       []string{"3.14", "2.71"},
			expected: true,
		},
		{
			name:     "float32 invalid value",
			val:      float32(1.5),
			of:       []string{"3.14", "2.71"},
			expected: false,
		},
		{
			name:     "float64 valid value",
			val:      3.14159,
			of:       []string{"3.14159", "2.71828"},
			expected: true,
		},
		{
			name:     "float64 invalid value",
			val:      1.23,
			of:       []string{"3.14", "2.71"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOneOfValidateString(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "string valid value",
			val:      "admin",
			of:       []string{"admin", "user", "guest"},
			expected: true,
		},
		{
			name:     "string invalid value",
			val:      "moderator",
			of:       []string{"admin", "user", "guest"},
			expected: false,
		},
		{
			name:     "string case sensitive",
			val:      "Admin",
			of:       []string{"admin", "user"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOneOfValidatePointer(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "pointer to int valid",
			val:      ptrInt(42),
			of:       []string{"42", "100"},
			expected: true,
		},
		{
			name:     "pointer to int invalid",
			val:      ptrInt(99),
			of:       []string{"42", "100"},
			expected: false,
		},
		{
			name:     "pointer to string valid",
			val:      ptrString("admin"),
			of:       []string{"admin", "user"},
			expected: true,
		},
		{
			name:     "nil pointer",
			val:      (*int)(nil),
			of:       []string{"42"},
			expected: false,
		},
		{
			name:     "double pointer",
			val:      ptrPtrInt(42),
			of:       []string{"42"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOneOfValidateInterface(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "interface with int",
			val:      any(42),
			of:       []string{"42"},
			expected: true,
		},
		{
			name:     "interface with string",
			val:      any("test"),
			of:       []string{"test"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOneOfValidateEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "empty allowed list",
			val:      42,
			of:       []string{},
			expected: true,
		},
		{
			name:     "single allowed value match",
			val:      "only",
			of:       []string{"only"},
			expected: true,
		},
		{
			name:     "single allowed value no match",
			val:      "other",
			of:       []string{"only"},
			expected: false,
		},
		{
			name:     "zero int",
			val:      0,
			of:       []string{"0"},
			expected: true,
		},
		{
			name:     "negative int",
			val:      -42,
			of:       []string{"-42", "100"},
			expected: true,
		},
		{
			name:     "empty string",
			val:      "",
			of:       []string{"", "value"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestOneOfValidateUnsupportedTypes(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		of       []string
		expected bool
	}{
		{
			name:     "bool type",
			val:      true,
			of:       []string{"true"},
			expected: false,
		},
		{
			name:     "slice type",
			val:      []int{1, 2, 3},
			of:       []string{"1"},
			expected: false,
		},
		{
			name:     "map type",
			val:      map[string]int{"key": 1},
			of:       []string{"key"},
			expected: false,
		},
		{
			name:     "struct type",
			val:      struct{}{},
			of:       []string{"value"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationOneOf(tt.val, tt.of)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Helper functions
func ptrInt(i int) *int {
	return &i
}

func ptrString(s string) *string {
	return &s
}

func ptrPtrInt(i int) **int {
	p := &i
	return &p
}
