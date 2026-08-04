package validator

import (
	"github.com/danilosciarra/guard/validator"
	"testing"
)

func TestEmailValidateValidAddresses(t *testing.T) {
	// Test email valide in vari formati

	tests := []struct {
		name     string
		val      any
		expected bool
	}{
		{
			name:     "simple valid email",
			val:      "test@example.com",
			expected: true,
		},
		{
			name:     "email with subdomain",
			val:      "user@mail.example.com",
			expected: true,
		},
		{
			name:     "email with plus",
			val:      "user+tag@example.com",
			expected: true,
		},
		{
			name:     "email with dots in local part",
			val:      "first.last@example.com",
			expected: true,
		},
		{
			name:     "email with numbers",
			val:      "user123@example123.com",
			expected: true,
		},
		{
			name:     "email with hyphen in domain",
			val:      "user@my-domain.com",
			expected: true,
		},
		{
			name:     "email with uppercase",
			val:      "USER@EXAMPLE.COM",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationEmail(tt.val)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEmailValidateInvalidAddresses(t *testing.T) {
	// Test email non valide

	tests := []struct {
		name     string
		val      any
		expected bool
	}{
		{
			name:     "missing @",
			val:      "userexample.com",
			expected: false,
		},
		{
			name:     "missing domain",
			val:      "user@",
			expected: false,
		},
		{
			name:     "missing local part",
			val:      "@example.com",
			expected: false,
		},
		{
			name:     "empty string",
			val:      "",
			expected: false,
		},
		{
			name:     "plain text",
			val:      "notanemail",
			expected: false,
		},
		{
			name:     "double @",
			val:      "user@@example.com",
			expected: false,
		},
		{
			name:     "spaces in email",
			val:      "user @example.com",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationEmail(tt.val)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEmailValidateWithDisplayName(t *testing.T) {
	// Test email con display name, ad esempio "John Doe <john@example.com>"
	// Il validatore deve restituire false perché non accetta display name, solo indirizzi email puri

	tests := []struct {
		name     string
		val      any
		expected bool
	}{
		{
			name:     "email with display name should be invalid",
			val:      "John Doe <john@example.com>",
			expected: false,
		},
		{
			name:     "email with quoted display name should be invalid",
			val:      `"John Doe" <john@example.com>`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationEmail(tt.val)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEmailValidateNonStringTypes(t *testing.T) {
	// Test con tipi di dati non stringa, ad esempio int, float, bool, slice, struct, nil
	// Il validatore deve restituire false perché si aspetta una stringa

	tests := []struct {
		name     string
		val      any
		expected bool
	}{
		{
			name:     "int type",
			val:      42,
			expected: false,
		},
		{
			name:     "float type",
			val:      3.14,
			expected: false,
		},
		{
			name:     "bool type",
			val:      true,
			expected: false,
		},
		{
			name:     "slice type",
			val:      []string{"user@example.com"},
			expected: false,
		},
		{
			name:     "struct type",
			val:      struct{}{},
			expected: false,
		},
		{
			name:     "nil value",
			val:      nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationEmail(tt.val)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEmailValidatePointer(t *testing.T) {
	// Test puntatori a stringhe

	validEmail := "user@example.com"
	invalidEmail := "notanemail"

	tests := []struct {
		name     string
		val      any
		expected bool
	}{
		{
			name:     "pointer to valid email string",
			val:      &validEmail,
			expected: false,
		},
		{
			name:     "pointer to invalid email string",
			val:      &invalidEmail,
			expected: false,
		},
		{
			name:     "nil string pointer",
			val:      (*string)(nil),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.NewValidationEmail(tt.val)
			result := v.Validate("", nil)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
