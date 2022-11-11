package main

import (
	"bytes"
	"fmt"
)

func BufferMain() {

	var b bytes.Buffer

	fmt.Printf("%v %T", b, b)

	b.WriteString("ABC")
	b.WriteString("DEF")

	fmt.Fprintf(&b, " A number: %d, a string: %v\n", 10, "bird")
	b.WriteString("DONE")

	fmt.Println(b.String())

}
