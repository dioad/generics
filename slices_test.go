package generics

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dioad/filter"
)

var (
	TestErrNotEven = fmt.Errorf("not even")
)

func TestApplyWithErrors(t *testing.T) {
	t.Run("Test Apply", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		errs := Apply(func(a int) error {
			if a%2 == 0 {
				return nil
			}
			return TestErrNotEven
		}, arr)

		compactedErrs := filter.FilterSlice(errs, func(a error) bool {
			return a != nil
		})

		if len(compactedErrs) != 3 {
			t.Errorf("Expected 3 errors, got %d", len(errs))
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

func TestZip(t *testing.T) {
	t.Run("Test Zip", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := []int{5, 4, 3, 2, 1}

		zipped, err := Zip(arr1, arr2)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		for i, v := range zipped {
			if v.A != arr1[i] || v.B != arr2[i] {
				t.Errorf("Expected (%d, %d), got (%d, %d)", arr1[i], arr2[i], v.A, v.B)
			}
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
	errs := Apply(func(a int) error {
		if a%2 == 0 {
			return nil
		}
		return errors.New("testErr")
	}, arr)

	if errs != nil {
		fmt.Println(errs)
	}

	// Output: [testErr <nil> testErr <nil> testErr]
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
