package main

import (
	"errors"
	"fmt"
	rand2 "math/rand"
	"time"
)

var rand float64

func init() {
	fmt.Println("Called init() on ", time.Now())
	rand2.Seed(time.Now().UnixNano())
	rand := rand2.Float64()
	fmt.Printf("Generated a random float64 value: %f", rand)
}

func main() {
	fmt.Println("Program started up on ", time.Now())
	fmt.Println(rand)

	fmt.Println("Generating 2nd random float")

	if f, err := randWithErr(); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("2nd random float:", f)
	}

}

func randWithErr() (float64, error) {
	rand := rand2.Float64()
	if rand > 0.5 {
		return rand, nil
	} else {
		return 0, errors.New("random number le 0.5 error")
	}
}
