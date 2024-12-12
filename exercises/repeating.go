package main

import (
	"strings"
)

func FindRepeatingWords(inputString string) string {
	// declare a map called words and initialize it
	// with an empty map
	var words = make(map[string]int)

	// split the input string into words
	// and store them in a slice called wordsList
	wordsList := strings.Fields(inputString)

	// create a map of the words and their frequency
	for _, word := range wordsList {
		// normalise the word by converting it to lowercase
		// and removing any punctuation
		word = strings.ToLower(word)
		word = strings.Trim(word, ",.!?")
		words[word]++
	}

	// create a slice to store the repeating words
	var repeatingWords []string

	for word, count := range words {
		if count > 1 {
			repeatingWords = append(repeatingWords, word)
		}
	}
	// join the repeating words into a single string
	// use comma as the separator
	// and return the string
	return strings.Join(repeatingWords, ", ")

}
