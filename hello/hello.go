package main

import (
	"errors"
	"fmt"
	"time"

	//"log"

	//"example.com/greetings"
	"math/rand"
	"strings"
)

func main() {

	fmt.Printf("Error: %s", generateError().Error())
}

func generateError() error {
	err := errors.New("Wrapped error")
	return err
}
func disjunction(s string) string {
	if s != "a" && s != "b" {
		return "invalid"
	} else {
		return "valid"
	}
}

func generateLongString() string {
	var longString string

	longString = "rule name: " + strings.Repeat("_long_rule_name_ ", 100)

	return longString
}
func testFunc() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if rnd := r.Intn(50); rnd >= 50 {
		fmt.Printf("%d is GE 50", rnd)
	} else {
		fmt.Printf("%d is LT 50", rnd)
	}
}
