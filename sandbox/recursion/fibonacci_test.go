package recursion

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {

	testCases := []struct {
		Index          int
		ExpectedResult int
	}{
		{
			Index:          0,
			ExpectedResult: 1,
		},
		{
			Index:          1,
			ExpectedResult: 1,
		},
		{
			Index:          2,
			ExpectedResult: 2,
		},
		{
			Index:          3,
			ExpectedResult: 3,
		},
		{
			Index:          4,
			ExpectedResult: 5,
		},
		{
			Index:          5,
			ExpectedResult: 8,
		},
		{
			Index:          6,
			ExpectedResult: 13,
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Index [%v] Fibonacci", tt.Index), func(t *testing.T) {
			ActualResult := fibonacci(tt.Index)
			if tt.ExpectedResult != ActualResult {
				t.Errorf("Failed test: expected [%v] got [%v]", tt.ExpectedResult, ActualResult)
			}
		})
	}

}
