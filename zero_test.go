package generics

import "testing"

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
