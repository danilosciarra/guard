package validator

import (
	validation2 "github.com/danilosciarra/guard/validator"
	"testing"
)

func TestNewValidationMax(t *testing.T) {
	type args struct {
		val  any
		vMax int
	}
	tests := []struct {
		name string
		args args
		want *validation2.Max
	}{
		{
			name: "TestNewValidationMax with int",
			args: args{
				val:  10,
				vMax: 100,
			},
		},
		{
			name: "TestNewValidationMax with string",
			args: args{
				val:  "test",
				vMax: 10,
			},
		},
		{
			name: "TestNewValidationMax with slice",
			args: args{
				val:  []int{1, 2, 3},
				vMax: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validation2.NewValidationMax(tt.args.val, tt.args.vMax); got == nil {
				t.Errorf("NewValidationMax() = %v, want non-nil", got)
			}
		})
	}
}

func TestMax_Validate(t *testing.T) {
	type args struct {
		path  string
		value any
	}
	tests := []struct {
		name string
		val  any
		max  int
		args args
		want bool
	}{
		// Test per numeri interi (int)
		{
			name: "int - valore uguale al massimo",
			val:  10,
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore minore del massimo",
			val:  5,
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore maggiore del massimo",
			val:  15,
			max:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "int - valore negativo minore del massimo",
			val:  -5,
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per int8
		{
			name: "int8 - valore valido",
			val:  int8(50),
			max:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int8 - valore maggiore del massimo",
			val:  int8(120),
			max:  100,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per int16
		{
			name: "int16 - valore valido",
			val:  int16(1000),
			max:  2000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per int32
		{
			name: "int32 - valore valido",
			val:  int32(50000),
			max:  100000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per int64
		{
			name: "int64 - valore valido",
			val:  int64(999999),
			max:  1000000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per numeri interi senza segno (uint)
		{
			name: "uint - valore valido",
			val:  uint(50),
			max:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "uint - valore maggiore del massimo",
			val:  uint(150),
			max:  100,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per uint8
		{
			name: "uint8 - valore valido",
			val:  uint8(200),
			max:  255,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per uint16
		{
			name: "uint16 - valore valido",
			val:  uint16(30000),
			max:  65000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per uint32
		{
			name: "uint32 - valore valido",
			val:  uint32(1000000),
			max:  2000000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per uint64
		{
			name: "uint64 - valore valido",
			val:  uint64(999999),
			max:  1000000,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per float32
		{
			name: "float32 - valore valido",
			val:  float32(50.5),
			max:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float32 - valore maggiore del massimo",
			val:  float32(150.5),
			max:  100,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "float32 - valore uguale al massimo",
			val:  float32(100.0),
			max:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per float64
		{
			name: "float64 - valore valido",
			val:  float64(99.99),
			max:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float64 - valore maggiore del massimo",
			val:  float64(100.01),
			max:  100,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per stringhe
		{
			name: "string - lunghezza uguale al massimo",
			val:  "1234567890",
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - lunghezza minore del massimo",
			val:  "test",
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - lunghezza maggiore del massimo",
			val:  "12345678901",
			max:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "string - stringa vuota",
			val:  "",
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "string - stringa con caratteri unicode",
			val:  "àèìòù",
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per slice
		{
			name: "slice - lunghezza uguale al massimo",
			val:  []int{1, 2, 3, 4, 5},
			max:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice - lunghezza minore del massimo",
			val:  []string{"a", "b"},
			max:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "slice - lunghezza maggiore del massimo",
			val:  []int{1, 2, 3, 4, 5, 6},
			max:  5,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "slice - slice vuoto",
			val:  []int{},
			max:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per array
		{
			name: "array - lunghezza uguale al massimo",
			val:  [3]int{1, 2, 3},
			max:  3,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "array - lunghezza minore del massimo",
			val:  [2]string{"a", "b"},
			max:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "array - lunghezza maggiore del massimo",
			val:  [6]int{1, 2, 3, 4, 5, 6},
			max:  5,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per map
		{
			name: "map - dimensione uguale al massimo",
			val:  map[string]int{"a": 1, "b": 2, "c": 3},
			max:  3,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "map - dimensione minore del massimo",
			val:  map[string]int{"a": 1},
			max:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "map - dimensione maggiore del massimo",
			val:  map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			max:  3,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "map - map vuota",
			val:  map[string]int{},
			max:  5,
			args: args{path: "field", value: nil},
			want: true,
		},
		// Test per puntatori
		{
			name: "pointer to int - valore valido",
			val:  func() *int { v := 50; return &v }(),
			max:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "pointer to int - valore maggiore del massimo",
			val:  func() *int { v := 150; return &v }(),
			max:  100,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "pointer to string - lunghezza valida",
			val:  func() *string { v := "test"; return &v }(),
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "nil pointer",
			val:  (*int)(nil),
			max:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per interfacce
		{
			name: "interface with int - valore valido",
			val:  func() any { var v any = 50; return v }(),
			max:  100,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "interface with string - lunghezza valida",
			val:  func() any { var v any = "test"; return v }(),
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "nil interface",
			val:  func() any { var v any; return v }(),
			max:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test per tipi non supportati
		{
			name: "struct - tipo non supportato",
			val:  struct{ Field string }{Field: "test"},
			max:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "bool - tipo non supportato",
			val:  true,
			max:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		{
			name: "channel - tipo non supportato",
			val:  make(chan int),
			max:  10,
			args: args{path: "field", value: nil},
			want: false,
		},
		// Test edge cases
		{
			name: "int - zero con massimo positivo",
			val:  0,
			max:  10,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - zero con massimo zero",
			val:  0,
			max:  0,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "int - valore negativo con massimo negativo",
			val:  -10,
			max:  -5,
			args: args{path: "field", value: nil},
			want: true,
		},
		{
			name: "float - valore molto piccolo",
			val:  0.0001,
			max:  1,
			args: args{path: "field", value: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation2.NewValidationMax(tt.val, tt.max)
			if got := v.Validate(tt.args.path, tt.args.value); got != tt.want {
				t.Errorf("Max.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
