package timing

import (
	"fmt"
	"time"
)

func Calculation() {
	var a int
	var b int
	a = 1
	for i := 0; i < 1000; i++ {
		a += i
		b += a + (1 + b)
	}
}

func TimeCalculation() {
	start := time.Now()
	Calculation()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("The calculation took %v", elapsed)
}
