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
	main1()
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
