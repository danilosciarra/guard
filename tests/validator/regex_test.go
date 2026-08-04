package validator

import (
	validation2 "github.com/danilosciarra/guard/validator"
	"testing"
)

func TestNewValidationRegex(t *testing.T) {
	type args struct {
		val    any
		vRegex string
	}
	tests := []struct {
		name string
		args args
		want *validation2.Regex
	}{
		{
			name: "TestNewValidationRegex with string and pattern",
			args: args{
				val:    "test123",
				vRegex: "^[a-z]+[0-9]+$",
			},
		},
		{
			name: "TestNewValidationRegex with email pattern",
			args: args{
				val:    "test@example.com",
				vRegex: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			},
		},
		{
			name: "TestNewValidationRegex with numeric pattern",
			args: args{
				val:    "12345",
				vRegex: "^[0-9]+$",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validation2.NewValidationRegex(tt.args.val, tt.args.vRegex); got == nil {
				t.Errorf("NewValidationRegex() = %v, want non-nil", got)
			}
		})
	}
}

func TestRegex_Validate(t *testing.T) {
	type args struct {
		path  string
		value any
	}
	tests := []struct {
		name  string
		val   any
		regex string
		args  args
		want  bool
	}{
		// Test per stringhe base
		{
			name:  "string - match pattern alfabetico",
			val:   "hello",
			regex: "^[a-z]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "string - no match pattern alfabetico",
			val:   "hello123",
			regex: "^[a-z]+$",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		{
			name:  "string - match pattern alfanumerico",
			val:   "test123",
			regex: "^[a-zA-Z0-9]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "string - no match pattern alfanumerico",
			val:   "test_123",
			regex: "^[a-zA-Z0-9]+$",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		// Test per email
		{
			name:  "email - pattern valido",
			val:   "user@example.com",
			regex: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			args:  args{path: "email", value: nil},
			want:  true,
		},
		{
			name:  "email - pattern invalido",
			val:   "invalid.email",
			regex: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			args:  args{path: "email", value: nil},
			want:  false,
		},
		{
			name:  "email - con punto e underscore",
			val:   "user.name_123@example.co.uk",
			regex: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			args:  args{path: "email", value: nil},
			want:  true,
		},
		// Test per numeri
		{
			name:  "numero - solo cifre",
			val:   "12345",
			regex: "^[0-9]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "numero - con lettere",
			val:   "123abc",
			regex: "^[0-9]+$",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		{
			name:  "numero intero convertito a string",
			val:   123,
			regex: "^[0-9]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		// Test per pattern specifici
		{
			name:  "codice postale italiano",
			val:   "00100",
			regex: "^[0-9]{5}$",
			args:  args{path: "cap", value: nil},
			want:  true,
		},
		{
			name:  "codice postale italiano - troppo corto",
			val:   "001",
			regex: "^[0-9]{5}$",
			args:  args{path: "cap", value: nil},
			want:  false,
		},
		{
			name:  "codice postale italiano - troppo lungo",
			val:   "001000",
			regex: "^[0-9]{5}$",
			args:  args{path: "cap", value: nil},
			want:  false,
		},
		{
			name:  "telefono - formato italiano",
			val:   "+39 123 4567890",
			regex: "^\\+39 [0-9]{3} [0-9]{7,10}$",
			args:  args{path: "phone", value: nil},
			want:  true,
		},
		{
			name:  "telefono - formato invalido",
			val:   "123456",
			regex: "^\\+39 [0-9]{3} [0-9]{7,10}$",
			args:  args{path: "phone", value: nil},
			want:  false,
		},
		// Test per URL
		{
			name:  "url - http valido",
			val:   "http://example.com",
			regex: "^https?://[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}",
			args:  args{path: "url", value: nil},
			want:  true,
		},
		{
			name:  "url - https valido",
			val:   "https://example.com",
			regex: "^https?://[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}",
			args:  args{path: "url", value: nil},
			want:  true,
		},
		{
			name:  "url - invalido",
			val:   "ftp://example.com",
			regex: "^https?://[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}",
			args:  args{path: "url", value: nil},
			want:  false,
		},
		// Test per stringhe vuote
		{
			name:  "stringa vuota - match con pattern che accetta vuoto",
			val:   "",
			regex: "^[a-z]*$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "stringa vuota - no match con pattern che richiede almeno un carattere",
			val:   "",
			regex: "^[a-z]+$",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		// Test per pattern con caratteri speciali
		{
			name:  "password - con maiuscola minuscola numero e speciale",
			val:   "Pass123!",
			regex: "^[a-zA-Z0-9!@#$%^&*]+$",
			args:  args{path: "password", value: nil},
			want:  true,
		},
		{
			name:  "password - senza caratteri speciali",
			val:   "Password123",
			regex: "^[a-zA-Z0-9!@#$%^&*]+$",
			args:  args{path: "password", value: nil},
			want:  true,
		},
		{
			name:  "password - lunghezza minima 8 caratteri",
			val:   "Pass123!",
			regex: "^.{8,}$",
			args:  args{path: "password", value: nil},
			want:  true,
		},
		{
			name:  "password - lunghezza minima non soddisfatta",
			val:   "Pass12!",
			regex: "^.{8,}$",
			args:  args{path: "password", value: nil},
			want:  false,
		},
		// Test con tipi non stringa
		{
			name:  "float convertito a string",
			val:   123.45,
			regex: "^[0-9]+\\.[0-9]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "bool convertito a string - true",
			val:   true,
			regex: "^(true|false)$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "bool convertito a string - false",
			val:   false,
			regex: "^(true|false)$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		// Test per pattern invalidi (che causano errore)
		{
			name:  "regex invalido - parentesi non chiusa",
			val:   "test",
			regex: "^[a-z",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		{
			name:  "regex invalido - pattern malformato",
			val:   "test",
			regex: "(?P<invalid",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		// Test per case sensitivity
		{
			name:  "case sensitive - maiuscole non matchano minuscole",
			val:   "HELLO",
			regex: "^[a-z]+$",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		{
			name:  "case insensitive - con flag (?i)",
			val:   "HELLO",
			regex: "(?i)^[a-z]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		// Test per pattern con spazi
		{
			name:  "stringa con spazi - match",
			val:   "hello world",
			regex: "^[a-z ]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "stringa con spazi - no match",
			val:   "hello world",
			regex: "^[a-z]+$",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		// Test per codice fiscale italiano
		{
			name:  "codice fiscale - formato valido",
			val:   "RSSMRA80A01H501U",
			regex: "^[A-Z]{6}[0-9]{2}[A-Z][0-9]{2}[A-Z][0-9]{3}[A-Z]$",
			args:  args{path: "cf", value: nil},
			want:  true,
		},
		{
			name:  "codice fiscale - formato invalido",
			val:   "RSSMRA80",
			regex: "^[A-Z]{6}[0-9]{2}[A-Z][0-9]{2}[A-Z][0-9]{3}[A-Z]$",
			args:  args{path: "cf", value: nil},
			want:  false,
		},
		// Test per partita IVA italiana
		{
			name:  "partita iva - formato valido",
			val:   "12345678901",
			regex: "^[0-9]{11}$",
			args:  args{path: "piva", value: nil},
			want:  true,
		},
		{
			name:  "partita iva - formato invalido",
			val:   "123456789",
			regex: "^[0-9]{11}$",
			args:  args{path: "piva", value: nil},
			want:  false,
		},
		// Test per IP address
		{
			name:  "ip address - formato semplice",
			val:   "192.168.1.1",
			regex: "^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",
			args:  args{path: "ip", value: nil},
			want:  true,
		},
		{
			name:  "ip address - formato invalido",
			val:   "192.168.1",
			regex: "^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",
			args:  args{path: "ip", value: nil},
			want:  false,
		},
		// Test per date
		{
			name:  "data - formato DD/MM/YYYY",
			val:   "31/12/2023",
			regex: "^[0-9]{2}/[0-9]{2}/[0-9]{4}$",
			args:  args{path: "date", value: nil},
			want:  true,
		},
		{
			name:  "data - formato YYYY-MM-DD",
			val:   "2023-12-31",
			regex: "^[0-9]{4}-[0-9]{2}-[0-9]{2}$",
			args:  args{path: "date", value: nil},
			want:  true,
		},
		// Test edge cases
		{
			name:  "stringa con caratteri unicode",
			val:   "café",
			regex: "^[a-zà-ù]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
		{
			name:  "stringa con emoji",
			val:   "hello😀",
			regex: "^[a-z]+$",
			args:  args{path: "field", value: nil},
			want:  false,
		},
		{
			name:  "numero negativo",
			val:   -123,
			regex: "^-?[0-9]+$",
			args:  args{path: "field", value: nil},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation2.NewValidationRegex(tt.val, tt.regex)
			if got := v.Validate(tt.args.path, tt.args.value); got != tt.want {
				t.Errorf("Regex.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
