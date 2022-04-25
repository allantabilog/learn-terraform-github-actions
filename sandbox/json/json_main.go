package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshallJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		log.Println("error unmarshalling: ", err)
		return err
	}

	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}
	return nil
}

func (a Animal) MarshallJSON() ([]byte, error) {
	var s string

	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}
	return json.Marshal(s)
}

func test1() {
	blob := `["gopher", "armadillo", "zebra", "unknown", "gopher", "bee", "gopher", "zebra"]`
	var zoo []Animal

	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras: %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])
}

func vertexTest() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Vertex.Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
