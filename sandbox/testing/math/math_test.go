package math

import "testing"

// func TestAdd(t *testing.T) {
// 	got := Add(4, 6)
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{2, -3, -1},
	addTest{4, 8, 12},
	addTest{-3, 3, 0},
	addTest{1, 1, 2},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %d not equalt to expected %d", output, test.expected)
		}
	}
}

func TestWeirdAdd(t *testing.T) {
	got := WeirdAdd(1, 1)
	want := 2 + 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func IgnoreTestWeirdAddTable(t *testing.T) {
	for _, test := range addTests {
		if sum := WeirdAdd(test.arg1, test.arg2); sum != test.expected {
			t.Errorf("Sum %d not equal to expected %d", sum, test.expected)
		}
	}
}

func BenchMarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}
