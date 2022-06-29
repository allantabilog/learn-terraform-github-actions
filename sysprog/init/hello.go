package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	now := time.Now()
	earlier := now.Add(-10 * time.Minute)

	fmt.Println(earlier.Before(now))

	file, err := os.Open("./hello.go")

	if err != nil {
		log.Fatal(err)
	}

	// Read the contents of the file into a slice of bytes
	data := make([]byte, 2000)
	count, err := file.Read(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes: %q\n", count, data[:count])

	fmt.Println("Reading file info")

	fileStat, err := os.Stat("./hello.go")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()

}
