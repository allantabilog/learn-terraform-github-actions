package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func Test_Strings(t *testing.T) {
	suffixTests := []struct {
		fullString string
		suffix     string
		expected   bool
	}{
		{"golang", "lang", true},
		{"golang", "ng", true},
		{"golang", "g", true},
		{"golang", "x", false},
	}
	prefixTests := []struct {
		fullString string
		prefix     string
		expected   bool
	}{
		{"golang", "gol", true},
		{"golang", "go", true},
		{"golang", "g", true},
		{"golang", "x", false},
	}

	for _, tc := range suffixTests {
		t.Run("suffix tests", func(t *testing.T) {
			assert.Equal(t, strings.HasSuffix(tc.fullString, tc.suffix), tc.expected)
		})
	}
	for _, tc := range prefixTests {
		t.Run("prefix tests", func(t *testing.T) {
			assert.Equal(t, strings.HasPrefix(tc.fullString, tc.prefix), tc.expected)
		})
	}
	t.Run("repeat string test", func(t *testing.T) {
		assert.Equal(t, "aaaaaaa", strings.Repeat("a", 7))
	})
	t.Run("string reader test", func(t *testing.T) {
		assert.True(t, true)

		reader := strings.NewReader("hello")
		byte, error := reader.ReadByte()
		if error != nil {
			log.Println("error while reading a byte")
		}
		log.Println("byte: ", byte)
	})

	t.Run("String builder test", func(t *testing.T) {
		var b strings.Builder
		for i := 3; i >= 1; i-- {
			fmt.Fprintf(&b, "%d...", i)
		}
		b.WriteString("ignition")
		fmt.Println(b.String())
	})
}
