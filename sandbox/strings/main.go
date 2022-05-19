package main

import (
	"fmt"
	"sort"
)

func main() {
	options := []string{
		"Transfer not created",
		"Mortgage not created",
		"Mortgage not signed",
		"Mortgage signed",
		"Discharge not created",
		"Discharge not signed",
		"Discharge signed",
	}

	sort.Strings(options)

	for _, option := range options {
		fmt.Println(option)
	}

}
