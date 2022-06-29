package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"sandbox/functions/misc"
	"strings"
)

func main() {

	fmt.Println(strings.EqualFold("abc", "ABC"))
	fmt.Println(strings.Contains("abc", "BC"))

}

func functionPointer() {
	inc := increment

	fmt.Println(inc(0))
	fmt.Println(inc(1))
	fmt.Println(inc(inc(inc(0))))
}
func increment(n int) int {
	return n + 1
}
func runtimeCaller() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()

	a := 2 * 5
	where()
	b := a + 100
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
}

func factoryTester() {
	//fmt.Println(addBmpSuffix("picture"))
	//fmt.Println(addJpgSuffix("picture"))

	bmpAdder := addSuffix(".bmp")
	jpgAdder := addSuffix(".jpg")

	fmt.Println(bmpAdder("picture1"))
	fmt.Println(bmpAdder("picture2"))
	fmt.Println(jpgAdder("picture"))
}

// a factory function
func addSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func addBmpSuffix(file string) string {
	return file + ".bmp"
}
func addJpgSuffix(file string) string {
	return file + ".jpg"
}

func tests3() {
	p2 := add2()
	fmt.Println(p2(3))
	p3 := adder(10)
	fmt.Println(p3(12))
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func funcName() {
	var evens []int

	evens = append(evens, 0)
	evens = append(evens, 2)
	fmt.Println(evens)
}

func tests2() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice: ", slice)
	odd := misc.Filter(slice, misc.IsOdd)
	fmt.Println("Odd elements: ", odd)
	even := misc.Filter(slice, misc.IsEven)
	fmt.Println("Even elements: ", even)

	fmt.Println(misc.IsEvenV2(100))
	fmt.Println(misc.IsEvenV2(101))
}

func typecast() {
	var badboys int = -1921
	var badboys2 = float64(badboys)
	var badboys3 = uint(badboys)

	fmt.Println(badboys2, badboys3)
}

func tripleCalls() {
	fmt.Println(triple(1, 2))
	fmt.Println(triple(1, -2))
	fmt.Println(triple(-1, 2))
	fmt.Println(triplev2(1, 2))
	fmt.Println(triplev2(1, -2))
	fmt.Println(triplev2(-1, 2))

}

func triple(a, b int) (sum, prod, diff int) {
	sum = a + b
	prod = a * b
	diff = a - b
	return
}

func triplev2(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func errorTest() {
	fmt.Println(Hello("world"))

	result, err := Hello("")
	if err != nil {
		fmt.Printf("Error calling Hello(): %v", err)
	}
	fmt.Println(result)
}

func div(n int, m int) (bool, error) {
	if m == 0 {

	}
	return false, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(fmt.Sprintf("[%v] is an empty name", name))
	}

	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message, nil
}
