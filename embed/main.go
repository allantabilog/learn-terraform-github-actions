package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
var f embed.FS

//embedding into an alias of FS
type myFile = embed.FS

//go:embed hello.txt
var ff myFile

//go:embed *
var devdir embed.FS

func main() {
	fmt.Printf("Embedded string: [%s]\n", s)
	fmt.Printf("Embedded byte array: [%v]\n", b)
	fmt.Printf("Embedded byte array: [%v]\n", string(b))

	contents, _ := f.ReadFile("hello.txt")
	fmt.Printf("Embedded file contents: [%v]\n", string(contents))

	contents, _ = ff.ReadFile("hello.txt")
	fmt.Printf("Embedded file contents: [%v]\n", string(contents))

	entries, _ := devdir.ReadDir(".")

	for _, e := range entries {
		fmt.Printf("File in dir: %v\n", e)
	}
	items := []string{"1", "2", "3"}
	fmt.Printf("contains 1: %v\n", contains(items, "1"))
	fmt.Printf("contains 4: %v\n", contains(items, "4"))
}

func contains(slice []string, str string) bool {
	for _, r := range slice {
		if str == r {
			return true
		}
	}
	return false
}
