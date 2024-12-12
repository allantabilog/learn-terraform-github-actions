package random

import (
	"fmt"
	"math/rand"
)


type RandomGenerator struct {
	prefix string
}

func NewRandomGenerator(prefix string) *RandomGenerator {
	return &RandomGenerator{prefix: prefix}
}

func (r *RandomGenerator) Generate() string {
	randomNumber := rand.Intn(1000)
	return fmt.Sprintf("%s%d", r.prefix, randomNumber)
}