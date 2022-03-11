package math

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		description     string
		n1              int
		n2              int
		expectedSum     int
		expectedProduct int
	}{
		{"basic test 1", 1, 2, 3, 2},
		{"basic test 2", 100, -2, 98, -200},
		{"basic test 3", -1, 1, 0, -1},
		{"basic test 4", -1, 11, 10, -11},
		{"basic test 5", -1, 1, 0, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expectedSum, Add(tc.n1, tc.n2))
			require.Equal(t, tc.expectedProduct, Multiply(tc.n1, tc.n2))
		})
	}
}

func TestSomething(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")

	var name *string
	assert.Equal(t, 1, 2, "this is expected to fail")
	assert.Nil(t, name)
	require.True(t, 1 == 2)
	assert.Equal(t, 1, 1, "this is expected to pass")

}
