package main

import (
	"fmt"

	"github.com/allantabilog/id-generator/data/random"
	"github.com/allantabilog/id-generator/data/uuid"
)

func main() {
	rg := random.NewRandomGenerator("order_")
	id := rg.Generate()
	fmt.Println(id)

	ug := uuid.NewUUIDGenerator("order_")
	id = ug.Generate()
	fmt.Println(id)
}