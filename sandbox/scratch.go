package main

import (
	"fmt"
	"strings"
)

func main(){
	BuildWordCountMap("hello world")
	BuildWordCountMap("hello hello there world")
	BuildWordCountMap("hello there there world")
	BuildWordCountMap("hello world world")
	BuildWordCountMap("hello hello there there world world world")
}
func BuildWordCountMap(inputString string){

	var words = strings.Fields(inputString)
	var wordCountMap = make(map[string]int)

	// iterate over words in words array
	for _, word := range words {
		_ , present := wordCountMap[word]

		if present {
			wordCountMap[word] = wordCountMap[word] + 1
		} else {
			wordCountMap[word] = 1
		}
	}

	fmt.Println(inputString, wordCountMap)

}