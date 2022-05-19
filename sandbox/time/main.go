package main

import (
	"fmt"
	"time"
)

func main() {
	tNow := time.Now().Add(-5 * time.Minute)

	tUnix := tNow.Unix()

	fmt.Println("Go time ", tNow)
	fmt.Println("Unix time ", tUnix)
}
