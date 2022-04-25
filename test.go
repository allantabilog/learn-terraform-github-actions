package main

// You can edit this code!
// Click here and start typing.

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const defaultWord = "testWord"

func main() {
	// get command line arg
	if len(os.Args) > 1 {
		word := os.Args[1]
		regexTest(word)
	} else {
		regexTest("nax")
		regexTest("Backy")
		regexTest("Pexa")
		regexTest("ta	st.pass")
	}

}
func regexTest(word string) {

	vowels := [5]string{"a", "e", "i", "o", "u"}
	fmt.Println()
	for _, vowel := range vowels {
		fmt.Print(" ", strings.ReplaceAll(word, "a", vowel))
	}

}
func regexTest1(word string) {
	fmt.Println(word)
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.MatchString("oeach"))
	fmt.Println(r.ReplaceAllString("a peach", "fruit"))

}
func regexTest0(word string) {
	fmt.Println(word)
	fmt.Println("Input word: ", word)
	// print vowel-substitutions of input word
	regex := regexp.MustCompile("a")
	vowels := [5]string{"a", "e", "i", "o", "u"}
	for _, v := range vowels {
		fmt.Println(v)
		fmt.Println(regex.ReplaceAllString(word, v))
	}
}
