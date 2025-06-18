package generics

import (
	"errors"
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "Filter even numbers",
			arr:       []int{1, 2, 3, 4, 5},
			predicate: func(a int) bool { return a%2 == 0 },
			expected:  []int{2, 4},
		},
		{
			name:      "Filter all",
			arr:       []int{1, 2, 3, 4, 5},
			predicate: func(a int) bool { return true },
			expected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "Filter none",
			arr:       []int{1, 2, 3, 4, 5},
			predicate: func(a int) bool { return false },
			expected:  []int{},
		},
		{
			name:      "Filter empty array",
			arr:       []int{},
			predicate: func(a int) bool { return a%2 == 0 },
			expected:  []int{},
		},
		{
			name:      "Filter nil array",
			arr:       nil,
			predicate: func(a int) bool { return a%2 == 0 },
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := Filter(tt.arr, tt.predicate)

			if len(filtered) != len(tt.expected) {
				t.Errorf("Expected length %v, got %v", len(tt.expected), len(filtered))
			}

			for i, v := range filtered {
				if v != tt.expected[i] {
					t.Errorf("Expected arr[%d]==%v, got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		initial  int
		reducer  func(int, int) int
		expected int
	}{
		{
			name:     "Sum",
			arr:      []int{1, 2, 3, 4, 5},
			initial:  0,
			reducer:  func(acc, val int) int { return acc + val },
			expected: 15,
		},
		{
			name:     "Product",
			arr:      []int{1, 2, 3, 4, 5},
			initial:  1,
			reducer:  func(acc, val int) int { return acc * val },
			expected: 120,
		},
		{
			name:    "Max",
			arr:     []int{1, 7, 3, 9, 5},
			initial: 0,
			reducer: func(acc, val int) int {
				if val > acc {
					return val
				} else {
					return acc
				}
			},
			expected: 9,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			initial:  42,
			reducer:  func(acc, val int) int { return acc + val },
			expected: 42,
		},
		{
			name:     "Nil array",
			arr:      nil,
			initial:  42,
			reducer:  func(acc, val int) int { return acc + val },
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reduce(tt.arr, tt.initial, tt.reducer)

			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "Contains even number",
			arr:       []int{1, 2, 3, 4, 5},
			predicate: func(a int) bool { return a%2 == 0 },
			expected:  true,
		},
		{
			name:      "Does not contain even number",
			arr:       []int{1, 3, 5, 7, 9},
			predicate: func(a int) bool { return a%2 == 0 },
			expected:  false,
		},
		{
			name:      "Empty array",
			arr:       []int{},
			predicate: func(a int) bool { return a%2 == 0 },
			expected:  false,
		},
		{
			name:      "Nil array",
			arr:       nil,
			predicate: func(a int) bool { return a%2 == 0 },
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.arr, tt.predicate)

			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	t.Run("No errors", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		sum := 0

		err := ForEach(arr, func(a int) error {
			sum += a
			return nil
		})

		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if sum != 15 {
			t.Errorf("Expected sum 15, got %v", sum)
		}
	})

	t.Run("With error", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		sum := 0
		expectedErr := errors.New("test error")

		err := ForEach(arr, func(a int) error {
			sum += a
			if a == 3 {
				return expectedErr
			}
			return nil
		})

		if err != expectedErr {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}

		if sum != 6 {
			t.Errorf("Expected sum 6 (1+2+3), got %v", sum)
		}
	})

	t.Run("Empty array", func(t *testing.T) {
		arr := []int{}
		called := false

		err := ForEach(arr, func(a int) error {
			called = true
			return nil
		})

		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if called {
			t.Errorf("Function should not have been called for empty array")
		}
	})

	t.Run("Nil array", func(t *testing.T) {
		var arr []int = nil
		called := false

		err := ForEach(arr, func(a int) error {
			called = true
			return nil
		})

		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if called {
			t.Errorf("Function should not have been called for nil array")
		}
	})
}

// Examples

func ExampleFilter() {
	arr := []int{1, 2, 3, 4, 5}
	filtered := Filter(arr, func(a int) bool {
		return a%2 == 0
	})

	fmt.Println(filtered)
	// Output: [2 4]
}

func ExampleReduce() {
	arr := []int{1, 2, 3, 4, 5}
	sum := Reduce(arr, 0, func(acc int, a int) int {
		return acc + a
	})

	fmt.Println(sum)
	// Output: 15
}

func ExampleContains() {
	arr := []int{1, 2, 3, 4, 5}
	hasEven := Contains(arr, func(a int) bool {
		return a%2 == 0
	})

	fmt.Println(hasEven)
	// Output: true
}

func ExampleForEach() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0

	err := ForEach(arr, func(a int) error {
		sum += a
		return nil
	})

	fmt.Println("Error:", err)
	fmt.Println("Sum:", sum)
	// Output:
	// Error: <nil>
	// Sum: 15
}
