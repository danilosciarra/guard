package tests

import (
	"encoding/json"
	"testing"

	"github.com/danilosciarra/guard"
)

func TestNewValidationManager(t *testing.T) {
	got := guard.New()
	if got == nil {
		t.Errorf("New() = %v, want non-nil Guard", got)
	}
}

func TestManager_Validate(t *testing.T) {
	type nestedObj struct {
		NestedField1 string `validate:"required"`
	}
	type testObj struct {
		Field1 string     `validate:"required,min=3"`
		Field2 int        `validate:"min=10,max=100"`
		Field3 string     `validate:"regex=^[a-zA-Z]+$"`
		Field4 string     `validate:"required"`
		Nested *nestedObj `validate:"required"`
		Array  []string   `validate:"min=1"`
	}

	var obj = testObj{
		Field1: "abc",
		Field2: 50,
		Field3: "TestString",
		Field4: "someValue",
		Nested: &nestedObj{NestedField1: "nestedValue"},
		Array:  []string{"item1"},
	}

	vm := guard.New()

	obj.Field1 = "ab"
	err := vm.Validate(&obj)
	if err == nil {
		t.Errorf("Expected validation error for Field1, got nil")
	} else {
		t.Log(err.Error())
	}

	obj.Field1 = "abc"
	obj.Field2 = 5
	err = vm.Validate(&obj)
	if err == nil {
		t.Errorf("Expected validation error for Field2, got nil")
	} else {
		t.Log(err.Error())
	}

	obj.Field2 = 50
	obj.Field3 = "123"
	err = vm.Validate(&obj)
	if err == nil {
		t.Errorf("Expected validation error for Field3, got nil")
	} else {
		t.Log(err.Error())
	}

	obj.Field3 = "ValidString"
	err = vm.Validate(&obj)
	if err != nil {
		t.Errorf("Validation failed: %v", err)
	}

	obj.Nested = &nestedObj{NestedField1: ""}
	vm = guard.New()
	err = vm.Validate(&obj)
	if err == nil {
		t.Errorf("Expected validation error for NestedField1, got nil")
		return
	}
	t.Log(err.Error())

	obj.Nested = &nestedObj{NestedField1: "nestedValue"}
	obj.Array = []string{}
	err = vm.Validate(&obj)
	if err == nil {
		t.Errorf("Expected validation error for Array, got nil")
	} else {
		t.Log(err.Error())
	}

	type myEnum string
	const (
		Value1 myEnum = "value1"
		Value2 myEnum = "value2"
		Value3 myEnum = "value3"
	)
	type testEnum struct {
		EnumField myEnum `validate:"regex=^(value1|value2|value3)$"`
	}
	enumObj := testEnum{EnumField: Value1}
	vmEnum := guard.New()
	err = vmEnum.Validate(&enumObj)
	if err != nil {
		t.Errorf("Validation failed for enum valid case: %v", err)
	}
	enumObj.EnumField = "invalidValue"
	err = vmEnum.Validate(&enumObj)
	if err == nil {
		t.Errorf("Expected validation error for enum invalid case, got nil")
	} else {
		t.Logf("Validation failed as expected: %v", err)
	}

	type testOneOf struct {
		OneOfField string `validate:"oneOf=option1|option2|option3"`
	}
	oneOfObj := testOneOf{OneOfField: "option1"}
	vmOneOf := guard.New()
	err = vmOneOf.Validate(&oneOfObj)
	if err != nil {
		t.Errorf("Validation failed for oneof valid case: %v", err)
	}
	oneOfObj.OneOfField = "invalidOption"
	err = vmOneOf.Validate(&oneOfObj)
	if err == nil {
		t.Errorf("Expected validation error for oneof invalid case, got nil")
	} else {
		t.Logf("Validation failed as expected: %v", err)
	}

	type testNode struct {
		NodeField string
	}
	type testWithoutTags struct {
		FieldA string
		FieldB int
		FieldC *testNode
	}
	objWithoutTags := testWithoutTags{FieldA: "anyValue", FieldB: 123}
	vmWithoutTags := guard.New()
	err = vmWithoutTags.Validate(&objWithoutTags)
	if err != nil {
		t.Errorf("Validation failed for struct without tags: %v", err)
	}
}

func TestSetCustomValidationFunc_Passes(t *testing.T) {
	type myStruct struct {
		Name string `validate:"custom_nonempty"`
	}
	obj := myStruct{Name: "hello"}
	vm := guard.New()

	vm.RegisterValidator("custom_nonempty", func(path string, value any) bool {
		return value != ""
	})

	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid, got error: %v", err)
	}
}

func TestSetCustomValidationFunc_Fails(t *testing.T) {
	type myStruct struct {
		Name string `validate:"custom_fail"`
	}
	obj := myStruct{Name: "something"}
	vm := guard.New()

	vm.RegisterValidator("custom_fail", func(path string, value any) bool {
		return false
	})

	err := vm.Validate(&obj)
	if err == nil {
		t.Error("Expected error from custom validator, got nil")
	}
}

func TestManager_ValidateSlice(t *testing.T) {
	type item struct {
		Name string `validate:"required"`
	}
	items := []item{
		{Name: "Alice"},
		{Name: "Bob"},
	}
	vm := guard.New()
	err := vm.Validate(items)
	if err != nil {
		t.Fatalf("Expected valid slice, got error: %v", err)
	}
}

func TestManager_ValidateSlice_InvalidItem(t *testing.T) {
	type item struct {
		Name string `validate:"required"`
	}
	items := []item{
		{Name: "Alice"},
		{Name: ""},
	}
	vm := guard.New()
	err := vm.Validate(items)
	if err == nil {
		t.Error("Expected error for invalid item in slice, got nil")
	}
}

// --- getValueByPath coverage ---

func TestManager_Validate_NestedPointerField(t *testing.T) {
	type Inner struct {
		Value string `validate:"required"`
	}
	type Outer struct {
		Inner *Inner `validate:"required"`
	}
	obj := Outer{Inner: &Inner{Value: "ok"}}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid nested pointer, got: %v", err)
	}
}

func TestManager_Validate_NestedPointerField_Invalid(t *testing.T) {
	type Inner struct {
		Value string `validate:"required"`
	}
	type Outer struct {
		Inner *Inner `validate:"required"`
	}
	obj := Outer{Inner: &Inner{Value: ""}}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err == nil {
		t.Error("Expected error for empty nested field")
	}
}

func TestManager_Validate_SliceField(t *testing.T) {
	type Row struct {
		Id   string `validate:"required"`
		Name string `validate:"required"`
	}
	type Parent struct {
		Items []Row
	}
	obj := Parent{Items: []Row{
		{Id: "1", Name: "Alice"},
		{Id: "2", Name: "Bob"},
	}}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid, got: %v", err)
	}
}

func TestManager_Validate_NilPointerField_Skipped(t *testing.T) {
	type Inner struct {
		Value string `validate:"required"`
	}
	type Outer struct {
		Inner *Inner
		Name  string `validate:"required"`
	}
	obj := Outer{Inner: nil, Name: "test"}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid when nil pointer has no validate tag, got: %v", err)
	}
}

// TestManager_Validate_SliceWithIntIndex exercises getValueByPath integer-index branch.
func TestManager_Validate_SliceWithIntIndex(t *testing.T) {
	type Item struct {
		Name string `validate:"required"`
	}
	type Parent struct {
		Items []Item
	}
	obj := Parent{Items: []Item{{Name: "a"}, {Name: "b"}}}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid, got: %v", err)
	}
}

// TestManager_Validate_SliceItemMissing exercises the invalid-index path in getValueByPath.
func TestManager_Validate_SliceItemInvalid(t *testing.T) {
	type Item struct {
		Name string `validate:"required"`
	}
	type Parent struct {
		Items []Item
	}
	obj := Parent{Items: []Item{{Name: "a"}, {Name: ""}}}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err == nil {
		t.Error("Expected validation failure for empty Name in slice item")
	}
}

// TestManager_Validate_NestedPtrField exercises pointer field handling in getValueByPath.
func TestManager_Validate_NestedPtrField_Valid(t *testing.T) {
	type Inner struct {
		Value string `validate:"required"`
	}
	type Outer struct {
		Inner *Inner
	}
	obj := Outer{Inner: &Inner{Value: "filled"}}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// TestManager_Validate_EmptyStruct exercises the path where object has no validate tags.
func TestManager_Validate_NoTags(t *testing.T) {
	type Simple struct {
		Name string
		Age  int
	}
	obj := Simple{Name: "", Age: 0}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid (no tags), got: %v", err)
	}
}

func TestManager_ValidateExtendedStruct(t *testing.T) {
	type Base struct {
		BaseField string `validate:"required"`
	}
	type Extended struct {
		Base
		ExtField string `validate:"required"`
	}
	obj := Extended{Base: Base{BaseField: "base"}, ExtField: "ext"}
	vm := guard.New()
	err := vm.Validate(&obj)
	if err != nil {
		t.Fatalf("Expected valid, got: %v", err)
	}
	sObj := `{"BaseField":"base","ExtField":"ext"}`
	var objE Extended
	if err = json.Unmarshal([]byte(sObj), &objE); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	vm = guard.New()
	err = vm.Validate(&objE)
	if err != nil {
		t.Fatalf("Expected valid after unmarshal, got: %v", err)
	}
}
