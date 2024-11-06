package uuid

import (
	"fmt"

	"github.com/google/uuid"
)

type UUIDGenerator struct {
	prefix string
}

func NewUUIDGenerator(prefix string) *UUIDGenerator {
	return &UUIDGenerator{prefix: prefix}
}

func (u *UUIDGenerator) Generate() string {
	return fmt.Sprintf("%s-%s", u.prefix, uuid.New().String())
}