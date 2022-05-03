package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	earlier := now.Add(-10 * time.Minute)

	fmt.Println(earlier.Before(now))
}
