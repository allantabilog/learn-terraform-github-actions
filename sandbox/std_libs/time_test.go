package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Day(), now.Month(), now.Year())

	now = time.Now().UTC()
	fmt.Println(now.Format("02 Jan 2006 15:04"))
	fmt.Println(now.Format(time.RFC822))
}
