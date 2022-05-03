package functional

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {

	t.Run("Absolute value of a positive number", func(t *testing.T) {
		got := math.Abs(1)
		if got != 1 {
			t.Errorf("Expected 1 but got %v", got)
		}
	})

	t.Run("Absolute value of a negative number", func(t *testing.T) {
		got := math.Abs(-1)
		if got != 1 {
			t.Errorf("Expected 1 but got %v", got)
		}
	})

	t.Run("Absolute value of zero", func(t *testing.T) {
		got := math.Abs(0)
		if got != 0 {
			t.Errorf("Expected 0 but got %v", got)
		}
	})

}

func TestAbsTable(t *testing.T) {
	absTests := []struct {
		Description string
		Argument    float64
		Expected    float64
	}{
		{
			Description: "Absolute value of a negative number",
			Argument:    -1,
			Expected:    1,
		},
		{
			Description: "Absolute value of a positive number",
			Argument:    1,
			Expected:    1,
		},
		{
			Description: "Absolute value of zero",
			Argument:    -1,
			Expected:    1,
		},
	}

	for _, absTest := range absTests {
		t.Run(absTest.Description, func(t *testing.T) {
			got := math.Abs(absTest.Argument)
			if got != absTest.Expected {
				t.Errorf("Test failed: expected %v, got %v", absTest.Expected, got)
			}
		})
	}
}

func TestAllEven(t *testing.T) {
	testCases := []struct {
		Description    string
		Input          []int
		ExpectedResult bool
	}{
		{
			Description:    "All numbers in slice are even",
			Input:          []int{2, 4, 6, 8, 10, 22, 202},
			ExpectedResult: true,
		},
		{
			Description:    "Not all numbers in slice are even",
			Input:          []int{2, 5, 6, 8, 10, 22, 202},
			ExpectedResult: false,
		},
		{
			Description:    "Empty slice",
			Input:          []int{},
			ExpectedResult: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Description, func(t *testing.T) {
			ActualResult := AllEven(tt.Input)
			if ActualResult != tt.ExpectedResult {
				t.Errorf("Failed test, expected [%v] got [%v]", tt.ExpectedResult, ActualResult)
			}
		})
	}
}
