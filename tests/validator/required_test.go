package validator

import (
	validation2 "github.com/danilosciarra/guard/validator"
	"testing"
)

func TestNewValidationRequired(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name string
		args args
		want *validation2.Required
	}{
		{
			name: "TestNewValidationRequired with string",
			args: args{
				val: "test",
			},
		},
		{
			name: "TestNewValidationRequired with int",
			args: args{
				val: 10,
			},
		},
		{
			name: "TestNewValidationRequired with nil",
			args: args{
				val: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validation2.NewValidationRequired(tt.args.val); got == nil {
				t.Errorf("NewValidationRequired() = %v, want non-nil", got)
			}
		})
	}
}

func TestRequired_Validate(t *testing.T) {
	type args struct {
		path  string
		value any
	}
	tests := []struct {
		name string
		val  any
		args args
		want bool
	}{
		// Test per nil
		{
			name: "nil value",
			val:  nil,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per stringhe
		{
			name: "string - non vuota",
			val:  "test",
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - vuota",
			val:  "",
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "string - con spazi",
			val:  "   ",
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - singolo carattere",
			val:  "a",
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per slice
		{
			name: "slice - non vuoto",
			val:  []int{1, 2, 3},
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice - vuoto",
			val:  []int{},
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "slice di stringhe - non vuoto",
			val:  []string{"a", "b"},
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice di stringhe - vuoto",
			val:  []string{},
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "slice - singolo elemento",
			val:  []int{1},
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per array
		{
			name: "array - non vuoto",
			val:  [3]int{1, 2, 3},
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "array - con lunghezza zero",
			val:  [0]int{},
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per map
		{
			name: "map - non vuota",
			val:  map[string]int{"a": 1, "b": 2},
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "map - vuota",
			val:  map[string]int{},
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "map - singolo elemento",
			val:  map[string]string{"key": "value"},
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per int
		{
			name: "int - valore positivo",
			val:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore negativo",
			val:  -10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - zero",
			val:  0,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per int8
		{
			name: "int8 - valore non zero",
			val:  int8(5),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int8 - zero",
			val:  int8(0),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per int16
		{
			name: "int16 - valore non zero",
			val:  int16(100),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int16 - zero",
			val:  int16(0),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per int32
		{
			name: "int32 - valore non zero",
			val:  int32(1000),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int32 - zero",
			val:  int32(0),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per int64
		{
			name: "int64 - valore non zero",
			val:  int64(10000),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int64 - zero",
			val:  int64(0),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per float32
		{
			name: "float32 - valore positivo",
			val:  float32(10.5),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float32 - valore negativo",
			val:  float32(-10.5),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float32 - zero",
			val:  float32(0.0),
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "float32 - valore molto piccolo",
			val:  float32(0.0001),
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per float64
		{
			name: "float64 - valore positivo",
			val:  float64(100.99),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float64 - valore negativo",
			val:  float64(-100.99),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float64 - zero",
			val:  float64(0.0),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per puntatori
		{
			name: "pointer to string - non vuota",
			val:  func() *string { v := "test"; return &v }(),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "pointer to string - vuota",
			val:  func() *string { v := ""; return &v }(),
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "pointer to int - non zero",
			val:  func() *int { v := 10; return &v }(),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "pointer to int - zero",
			val:  func() *int { v := 0; return &v }(),
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "nil pointer to string",
			val:  (*string)(nil),
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "nil pointer to int",
			val:  (*int)(nil),
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "pointer to slice - non vuoto",
			val:  func() *[]int { v := []int{1, 2}; return &v }(),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "pointer to slice - vuoto",
			val:  func() *[]int { v := []int{}; return &v }(),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per puntatori multipli
		{
			name: "double pointer to string - non vuota",
			val:  func() **string { v := "test"; p := &v; return &p }(),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "double pointer to int - zero",
			val:  func() **int { v := 0; p := &v; return &p }(),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per interfacce
		{
			name: "interface with string - non vuota",
			val:  func() any { var v any = "test"; return v }(),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "interface with string - vuota",
			val:  func() any { var v any = ""; return v }(),
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "interface with int - non zero",
			val:  func() any { var v any = 10; return v }(),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "interface with int - zero",
			val:  func() any { var v any = 0; return v }(),
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "nil interface",
			val:  func() any { var v any; return v }(),
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per tipi che ritornano sempre true (default case)
		{
			name: "bool - true",
			val:  true,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "bool - false",
			val:  false,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "struct - non nil",
			val:  struct{ Field string }{Field: "test"},
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "struct vuoto",
			val:  struct{}{},
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "channel",
			val:  make(chan int),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "func",
			val:  func() {},
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per uint types (non gestiti esplicitamente, quindi default case)
		{
			name: "uint - valore non zero",
			val:  uint(10),
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "uint - zero",
			val:  uint(0),
			args: args{path: "field", value: nil},
			want: true, // uint non è gestito esplicitamente, quindi default case restituisce true
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation2.NewValidationRequired(tt.val)
			if got := v.Validate(tt.args.path, tt.args.value); got != tt.want {
				t.Errorf("Required.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestRequired_Float covers the reflect.Float32/64 branch.
func TestRequired_Float(t *testing.T) {
	// non-zero float → required = true
	v := validation2.NewValidationRequired(float64(3.14))
	if !v.Validate("", nil) {
		t.Error("Expected true for non-zero float64")
	}
	// zero float → required = false
	v2 := validation2.NewValidationRequired(float64(0))
	if v2.Validate("", nil) {
		t.Error("Expected false for zero float64")
	}
	// float32
	v3 := validation2.NewValidationRequired(float32(1.0))
	if !v3.Validate("", nil) {
		t.Error("Expected true for non-zero float32")
	}
}

// TestRequired_Func covers the reflect.Func branch.
func TestRequired_Func(t *testing.T) {
	fn := func() {}
	v := validation2.NewValidationRequired(fn)
	if !v.Validate("", nil) {
		t.Error("Expected true for non-nil func")
	}
}

// TestRequired_Default covers the default branch (e.g. struct).
func TestRequired_DefaultType(t *testing.T) {
	type myStruct struct{ X int }
	v := validation2.NewValidationRequired(myStruct{X: 1})
	if !v.Validate("", nil) {
		t.Error("Expected true for non-nil struct (default branch)")
	}
}
