package generics

import (
	"errors"
	"fmt"
	"testing"
)

var (
	TestErrNotEven = fmt.Errorf("not even")
)

func TestApplyWithErrors(t *testing.T) {
	t.Run("Test Apply", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		err := Apply(func(a int) error {
			if a%2 == 0 {
				return nil
			}
			return TestErrNotEven
		}, arr)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		errCount := len(err.(*MapError).Errors)

		if errCount != 3 {
			t.Errorf("Expected 3 errors, got %d", errCount)
		}
	})
}

func TestApply(t *testing.T) {
	t.Run("Test Apply", func(t *testing.T) {
		expectedSum := 15
		runningSum := 0
		arr := []int{1, 2, 3, 4, 5}
		errs := Apply(func(a int) error {
			runningSum += a
			return nil
		}, arr)

		if errs != nil {
			t.Errorf("Expected nil, got %v", errs)
		}

		if runningSum != expectedSum {
			t.Errorf("Expected %d, got %d", expectedSum, runningSum)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("Test Map", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		doubled, err := Map(func(a int) (int, error) {
			return a * 2, nil
		}, arr)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		for i, v := range doubled {
			if v != arr[i]*2 {
				t.Errorf("Expected %d, got %d", arr[i]*2, v)
			}
		}
	})
}

func TestMapWithError(t *testing.T) {
	t.Run("Test Map with Error", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		doubled, err := Map(func(a int) (int, error) {
			if a%2 == 0 {
				return a * 2, nil
			}
			return 0, TestErrNotEven
		}, arr)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		mapError := err.(*MapError)
		errCount := len(mapError.Errors)

		if errCount != 3 {
			t.Errorf("Expected 3 errors, got %d", errCount)
		}

		for i, v := range doubled {
			if _, exists := mapError.Errors[i]; !exists {
				if v != arr[i]*2 {
					t.Errorf("Expected %d, got %d", arr[i]*2, v)
				}
			}
		}
	})
}

func TestSafeMap(t *testing.T) {
	t.Run("Test SafeMap", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		doubled := SafeMap(func(a int) int {
			return a * 2
		}, arr)

		for i, v := range doubled {
			if v != arr[i]*2 {
				t.Errorf("Expected %d, got %d", arr[i]*2, v)
			}
		}
	})
}

type ZipTest[A comparable, B comparable] struct {
	name        string
	arr1        []A
	arr2        []B
	expected    []Pair[A, B]
	expectedErr error
}

func zipTestHelper[A comparable, B comparable](t *testing.T, tests []ZipTest[A, B]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zipped, err := Zip(tt.arr1, tt.arr2)

			if err != nil {
				if tt.expectedErr == nil {
					t.Errorf("Expected nil, got %v", err)
				} else if err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected %v, got %v", tt.expectedErr, err)
				}
			}

			if len(zipped) != len(tt.expected) {
				t.Errorf("Expected %v, got %v", len(tt.expected), len(zipped))
			}

			for i, v := range zipped {
				if v.A != tt.expected[i].A || v.B != tt.expected[i].B {
					t.Errorf("Expected (%v, %v), got (%v, %v)", tt.expected[i].A, tt.expected[i].B, v.A, v.B)
				}
			}
		})
	}
}

func TestZipInt(t *testing.T) {
	tests := []ZipTest[int, int]{
		{
			name:        "Test Zip with Equal Lengths",
			arr1:        []int{1, 2, 3, 4, 5},
			arr2:        []int{5, 4, 3, 2, 1},
			expected:    []Pair[int, int]{{1, 5}, {2, 4}, {3, 3}, {4, 2}, {5, 1}},
			expectedErr: nil,
		},
		{
			name:        "Test Zip with Different Lengths",
			arr1:        []int{1, 2, 3, 4, 5},
			arr2:        []int{5, 4, 3, 2},
			expected:    nil,
			expectedErr: fmt.Errorf("arrays must be of equal length"),
		},
	}

	zipTestHelper(t, tests)
}

type CompactTest[T comparable] struct {
	name     string
	arr      []T
	expected []T
}

func compactTestHelper[A comparable](t *testing.T, tests []CompactTest[A]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compacted := Compact(tt.arr)

			if len(compacted) != len(tt.expected) {
				t.Errorf("Expected %v, got %v", len(tt.expected), len(compacted))
			}

			for i, v := range compacted {
				if v != tt.expected[i] {
					t.Errorf("Expected arr[%d]==%v, got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestCompactWithInt(t *testing.T) {
	tests := []CompactTest[int]{
		{
			name:     "Test Compact with No Zeros",
			arr:      []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test Compact with Zeros",
			arr:      []int{1, 2, 0, 0, 5},
			expected: []int{1, 2, 5},
		},
		{
			name:     "Test Compact with Empty Array",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "Test Compact with All Zeros",
			arr:      []int{0, 0, 0, 0, 0},
			expected: []int{},
		},
		{
			name:     "Test Compact with Nil",
			arr:      nil,
			expected: nil,
		},
	}

	compactTestHelper(t, tests)
}

func TestCompactWithErrors(t *testing.T) {
	tests := []CompactTest[error]{
		{
			name:     "Test Compact with No Zeros",
			arr:      []error{ErrDifferentLength, nil, nil},
			expected: []error{ErrDifferentLength},
		},
	}

	compactTestHelper[error](t, tests)
}

func TestSelectOne(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		selected, err := SelectOne(arr, func(a int) bool {
			return a%2 == 0
		})

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if selected != 2 {
			t.Errorf("Expected 2, got %d", selected)
		}
	})

	t.Run("not found", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		_, err := SelectOne(arr, func(a int) bool {
			return a%6 == 0
		})

		if !errors.Is(err, ErrNotFound) {
			t.Errorf("Expected ErrNotFound, got %v", err)
		}
	})
}

func ExampleMap() {
	arr := []int{1, 2, 3, 4, 5}
	doubled, err := Map(func(a int) (int, error) {
		return a * 2, nil
	}, arr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(doubled)
	// Output: [2 4 6 8 10]
}

func ExampleApply() {
	arr := []int{1, 2, 3, 4, 5}
	err := Apply(func(a int) error {
		if a%2 == 0 {
			return nil
		}
		return errors.New("testErr")
	}, arr)

	if err != nil {
		fmt.Println(err)
	}

	// Output: 3 errors
}

func ExampleZip() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 4, 3, 2, 1}

	zipped, err := Zip(arr1, arr2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(zipped)
	// Output: [{1 5} {2 4} {3 3} {4 2} {5 1}]
}

func ExampleCompact() {
	arr := []int{1, 2, 3, 4, 5}
	compacted := Compact(arr)

	fmt.Println(compacted)
	// Output: [1 2 3 4 5]
}

func ExampleSelectOne() {
	arr := []int{1, 2, 3, 4, 5}
	selected, err := SelectOne(arr, func(a int) bool {
		return a%2 == 0
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(selected)
	// Output:
	// 2
}
