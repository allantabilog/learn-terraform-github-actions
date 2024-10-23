package main

import (
	"fmt"
	"log"
	//"sandbox/maps/sandbox"
)

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60}

func main() {
	//sandbox.Main()
	main2()
}

func main2() {
	var counters = make(map[string]int, 10)
	counters["a"] = 1
	counters["b"] = 2
	counters["c"] = 3
	fmt.Printf("counters: %v\n", counters)

	modelToMake := map[string]string{
		"prius":    "toyota",
		"camry":    "toyota",
		"chevelle": "chevrolet",
		"mustang":  "ford",
	}

	fmt.Printf("modelToMake: %v\n", modelToMake)
	fmt.Println("prius==>", modelToMake["prius"])
	fmt.Println("lexus==>", modelToMake["lexus"])

	var make string = "tesla"
	if carMake, ok := modelToMake[make]; ok {
		fmt.Println("carMake==>", carMake)
	} else {
		fmt.Printf("carMake not found: %v", make)
	}

	for key, val := range modelToMake {
		fmt.Printf("key: %v, val: %v\n", key, val)
	}

}

func main1() {
	fmt.Println("testing the comma, ok pattern")

	offset := timeZone["EST"]
	fmt.Println(offset)

	ast, ok := timeZone["AST"]
	fmt.Println(ast, ok)

	seconds := Offset("AST")
	fmt.Println(seconds)
	seconds = Offset("EST")
	fmt.Println(seconds)

}

func Offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("Unknown timezone")
	return 0
}

func main0() {
	emails := make(map[string]string)

	emails["Bob"] = "The.Builder@gmail.com"
	emails["Bob2"] = "Another.Builder@gmail.com"
	emails["Bobby"] = "Yet.Another.Builder@gmail.com"

	fmt.Println(emails)
}
