package main

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"log"
)

func main() {
	resp, err := grab.Get(".", "http://www.golang-book.com/public/pdf/gobook.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to ", resp.Filename)
}
