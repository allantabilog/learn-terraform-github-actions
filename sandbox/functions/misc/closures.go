package misc

import "fmt"

func ClosureTest() {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf("- g is of type %T and has value %v\n", g, g)
	}
}

func GenInc(n int) func(x int) int {
	return func(x int) int {
		return x + n
	}
}

func Fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
