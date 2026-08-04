package validator

import (
	validation2 "github.com/danilosciarra/guard/validator"
	"testing"
)

func TestNewValidationMin(t *testing.T) {
	type args struct {
		val  any
		vMin int
	}
	tests := []struct {
		name string
		args args
		want *validation2.Min
	}{
		{
			name: "TestNewValidationMin with int",
			args: args{
				val:  10,
				vMin: 5,
			},
		},
		{
			name: "TestNewValidationMin with string",
			args: args{
				val:  "test",
				vMin: 1,
			},
		},
		{
			name: "TestNewValidationMin with slice",
			args: args{
				val:  []int{1, 2, 3},
				vMin: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validation2.NewValidationMin(tt.args.val, tt.args.vMin); got == nil {
				t.Errorf("NewValidationMin() = %v, want non-nil", got)
			}
		})
	}
}

func TestMin_Validate(t *testing.T) {
	type args struct {
		path  string
		value any
	}
	tests := []struct {
		name string
		val  any
		min  int
		args args
		want bool
	}{
		// Test per numeri interi (int)
		{
			name: "int - valore uguale al minimo",
			val:  10,
			min:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore maggiore del minimo",
			val:  15,
			min:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore minore del minimo",
			val:  5,
			min:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "int - valore negativo maggiore del minimo negativo",
			val:  -5,
			min:  -10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore negativo minore del minimo negativo",
			val:  -15,
			min:  -10,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per int8
		{
			name: "int8 - valore valido",
			val:  int8(100),
			min:  50,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int8 - valore minore del minimo",
			val:  int8(20),
			min:  50,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per int16
		{
			name: "int16 - valore valido",
			val:  int16(2000),
			min:  1000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per int32
		{
			name: "int32 - valore valido",
			val:  int32(100000),
			min:  50000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per int64
		{
			name: "int64 - valore valido",
			val:  int64(1000000),
			min:  999999,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per numeri interi senza segno (uint)
		{
			name: "uint - valore valido",
			val:  uint(100),
			min:  50,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "uint - valore minore del minimo",
			val:  uint(30),
			min:  50,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per uint8
		{
			name: "uint8 - valore valido",
			val:  uint8(200),
			min:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per uint16
		{
			name: "uint16 - valore valido",
			val:  uint16(50000),
			min:  30000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per uint32
		{
			name: "uint32 - valore valido",
			val:  uint32(2000000),
			min:  1000000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per uint64
		{
			name: "uint64 - valore valido",
			val:  uint64(1000000),
			min:  999999,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per float32
		{
			name: "float32 - valore valido",
			val:  float32(100.5),
			min:  50,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float32 - valore minore del minimo",
			val:  float32(30.5),
			min:  50,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "float32 - valore uguale al minimo",
			val:  float32(50.0),
			min:  50,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per float64
		{
			name: "float64 - valore valido",
			val:  float64(100.01),
			min:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float64 - valore minore del minimo",
			val:  float64(99.99),
			min:  100,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per stringhe
		{
			name: "string - lunghezza uguale al minimo",
			val:  "1234567890",
			min:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - lunghezza maggiore del minimo",
			val:  "12345678901",
			min:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - lunghezza minore del minimo",
			val:  "test",
			min:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "string - stringa vuota con minimo zero",
			val:  "",
			min:  0,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - stringa vuota con minimo positivo",
			val:  "",
			min:  1,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "string - stringa con caratteri unicode",
			val:  "àèìòùàèìòù",
			min:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per slice
		{
			name: "slice - lunghezza uguale al minimo",
			val:  []int{1, 2, 3, 4, 5},
			min:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice - lunghezza maggiore del minimo",
			val:  []string{"a", "b", "c", "d", "e", "f"},
			min:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice - lunghezza minore del minimo",
			val:  []int{1, 2, 3},
			min:  5,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "slice - slice vuoto con minimo zero",
			val:  []int{},
			min:  0,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice - slice vuoto con minimo positivo",
			val:  []int{},
			min:  1,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per array
		{
			name: "array - lunghezza uguale al minimo",
			val:  [3]int{1, 2, 3},
			min:  3,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "array - lunghezza maggiore del minimo",
			val:  [5]string{"a", "b", "c", "d", "e"},
			min:  2,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "array - lunghezza minore del minimo",
			val:  [2]int{1, 2},
			min:  5,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per map
		{
			name: "map - dimensione uguale al minimo",
			val:  map[string]int{"a": 1, "b": 2, "c": 3},
			min:  3,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "map - dimensione maggiore del minimo",
			val:  map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			min:  3,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "map - dimensione minore del minimo",
			val:  map[string]int{"a": 1},
			min:  3,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "map - map vuota con minimo zero",
			val:  map[string]int{},
			min:  0,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "map - map vuota con minimo positivo",
			val:  map[string]int{},
			min:  1,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per puntatori
		{
			name: "pointer to int - valore valido",
			val:  func() *int { v := 100; return &v }(),
			min:  50,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "pointer to int - valore minore del minimo",
			val:  func() *int { v := 30; return &v }(),
			min:  50,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "pointer to string - lunghezza valida",
			val:  func() *string { v := "test string"; return &v }(),
			min:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "nil pointer",
			val:  (*int)(nil),
			min:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per interfacce
		{
			name: "interface with int - valore valido",
			val:  func() any { var v any = 100; return v }(),
			min:  50,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "interface with string - lunghezza valida",
			val:  func() any { var v any = "test string"; return v }(),
			min:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "nil interface",
			val:  func() any { var v any; return v }(),
			min:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per tipi non supportati
		{
			name: "struct - tipo non supportato",
			val:  struct{ Field string }{Field: "test"},
			min:  1,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "bool - tipo non supportato",
			val:  true,
			min:  0,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "channel - tipo non supportato",
			val:  make(chan int),
			min:  0,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test edge cases
		{
			name: "int - zero con minimo zero",
			val:  0,
			min:  0,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - zero con minimo negativo",
			val:  0,
			min:  -10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore positivo con minimo negativo",
			val:  10,
			min:  -5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float - valore molto grande",
			val:  9999.9999,
			min:  1,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - lunghezza 1 con minimo 1",
			val:  "a",
			min:  1,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice - lunghezza 1 con minimo 1",
			val:  []int{1},
			min:  1,
			args: args{path: "field", value: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation2.NewValidationMin(tt.val, tt.min)
			if got := v.Validate(tt.args.path, tt.args.value); got != tt.want {
				t.Errorf("Min.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMin_Uint covers the reflect.Uint branch in Min.IsValid.
func TestMin_Uint(t *testing.T) {
	tests := []struct {
		name string
		val  uint
		min  int
		want bool
	}{
		{"uint above min", uint(10), 5, true},
		{"uint below min", uint(3), 5, false},
		{"uint equal min", uint(5), 5, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := validation2.NewValidationMin(tc.val, tc.min)
			if got := v.Validate("", nil); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// TestMin_Float covers the reflect.Float64 branch in Min.IsValid.
func TestMin_Float(t *testing.T) {
	tests := []struct {
		name string
		val  float64
		min  int
		want bool
	}{
		{"float above min", 5.5, 5, true},
		{"float below min", 4.9, 5, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := validation2.NewValidationMin(tc.val, tc.min)
			if got := v.Validate("", nil); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// TestMin_Map covers the reflect.Map branch in Min.IsValid.
func TestMin_Map(t *testing.T) {
	v := validation2.NewValidationMin(map[string]int{"a": 1, "b": 2}, 2)
	if !v.Validate("", nil) {
		t.Error("Expected true for map with len >= min")
	}
	v2 := validation2.NewValidationMin(map[string]int{"a": 1}, 2)
	if v2.Validate("", nil) {
		t.Error("Expected false for map with len < min")
	}
}

// TestMin_Default covers the default branch (unsupported type) in Min.IsValid.
func TestMin_DefaultType(t *testing.T) {
	v := validation2.NewValidationMin(true, 1) // bool is unsupported
	if v.Validate("", nil) {
		t.Error("Expected false for unsupported type bool")
	}
}
