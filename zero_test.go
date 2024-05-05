package generics

import (
	"fmt"
	"testing"
)

func TestIsZeroValue(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if !IsZeroValue(0) {
			t.Errorf("Expected true, got false")
		}

		if IsZeroValue(1) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("string", func(t *testing.T) {
		if !IsZeroValue("") {
			t.Errorf("Expected true, got false")
		}

		if IsZeroValue("a") {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("pointer", func(t *testing.T) {
		var a *int
		if !IsZeroValue(a) {
			t.Errorf("Expected true, got false")
		}

		b := 4
		if IsZeroValue(&b) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("struct", func(t *testing.T) {
		type testStruct struct {
			A int
			B string
		}

		var zeroStruct testStruct
		if !IsZeroValue(zeroStruct) {
			t.Errorf("Expected true, got false")
		}

		nonZeroStruct := testStruct{A: 1, B: "a"}
		if IsZeroValue(nonZeroStruct) {
			t.Errorf("Expected false, got true")
		}
	})
}

func ExampleIsZeroValue_int() {
	var a int
	fmt.Println(IsZeroValue(a))

	a = 4
	fmt.Println(IsZeroValue(a))
	// Output:
	// true
	// false
}

func ExampleIsZeroValue_struct() {
	type testStruct struct {
		A int
		B string
	}

	var b testStruct
	fmt.Println(IsZeroValue(b))

	b.A = 12
	b.B = "34"
	fmt.Println(IsZeroValue(b))

	// Output:
	// true
	// false
}

func ExampleIsZeroValue_string() {
	var a string
	fmt.Println(IsZeroValue(a))

	a = "asdf"
	fmt.Println(IsZeroValue(a))
	// Output:
	// true
	// false
}
