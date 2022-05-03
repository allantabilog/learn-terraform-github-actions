package closures

import (
	"fmt"
	"log"
	"runtime"
)

func closuresMain() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	a := 2 * 5
	where()
	b := a + 25
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
}
