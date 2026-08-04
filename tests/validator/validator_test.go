package validator

import (
	"github.com/danilosciarra/guard"
	val "github.com/danilosciarra/guard/validator"
	"reflect"
	"testing"
)

func TestType_Constants(t *testing.T) {
	tests := []struct {
		name     string
		typeVal  val.Type
		expected string
	}{
		{
			name:     "TypeMin constant",
			typeVal:  val.TypeMin,
			expected: "min",
		},
		{
			name:     "TypeMax constant",
			typeVal:  val.TypeMax,
			expected: "max",
		},
		{
			name:     "TypeRequired constant",
			typeVal:  val.TypeRequired,
			expected: "required",
		},
		{
			name:     "TypeRegex constant",
			typeVal:  val.TypeRegex,
			expected: "regex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.typeVal) != tt.expected {
				t.Errorf("Type constant = %v, want %v", tt.typeVal, tt.expected)
			}
		})
	}
}

func Test_validatorInterface(t *testing.T) {
	// Test che tutti i validatori implementino l'interfaccia validator
	tests := []struct {
		name      string
		validator any
	}{
		{
			name:      "Min implements validator",
			validator: val.NewValidationMin(10, 5),
		},
		{
			name:      "Max implements validator",
			validator: val.NewValidationMax(10, 100),
		},
		{
			name:      "Required implements validator",
			validator: val.NewValidationRequired("test"),
		},
		{
			name:      "Regex implements validator",
			validator: val.NewValidationRegex("test", "^[a-z]+$"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verifica che il validator abbia il metodo Validate
			v := reflect.ValueOf(tt.validator)
			method := v.MethodByName("Validate")
			if !method.IsValid() {
				t.Errorf("%s does not implement Validate method", tt.name)
				return
			}
			// Verifica che il metodo Validate abbia la firma corretta: (string, any) bool
			methodType := method.Type()
			if methodType.NumIn() != 2 || methodType.NumOut() != 1 {
				t.Errorf("%s Validate method has incorrect signature", tt.name)
			}
			if methodType.In(0).Kind() != reflect.String {
				t.Errorf("%s Validate method first parameter should be string", tt.name)
			}
			if methodType.Out(0).Kind() != reflect.Bool {
				t.Errorf("%s Validate method should return bool", tt.name)
			}
		})
	}
}

// TestManager_UnknownTag covers the default branch of newValidator (returns nil).
func TestManager_UnknownTag(t *testing.T) {
	type myStruct struct {
		Name string `validate:"unknowntag"`
	}
	obj := myStruct{Name: "hello"}
	g := guard.New()
	if g == nil {
		t.Error("Expected non-nil Guard")
		return
	}
	err := g.Validate(&obj)
	if err == nil {
		t.Fatalf("Expected error during validation for unknown tag, got nil")
	}
}

func TestManager_SliceFieldByIndex(t *testing.T) {
	type Row struct {
		Name string `validate:"required"`
	}
	type Parent struct {
		Items []Row
	}
	obj := Parent{Items: []Row{{Name: "a"}, {Name: "b"}}}
	g := guard.New()
	err := g.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid, got: %v", err)
	}
}

func TestManager_SliceFieldByStringId(t *testing.T) {
	type Row struct {
		Id   string `validate:"required"`
		Name string `validate:"required"`
	}
	type Parent struct {
		Items []Row
	}
	obj := Parent{Items: []Row{
		{Id: "row1", Name: "Alice"},
		{Id: "row2", Name: "Bob"},
	}}
	g := guard.New()
	err := g.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid, got: %v", err)
	}
}
