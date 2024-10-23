package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	for {
		each_record, err := reader.Read()
		if err != nil || err == io.EOF {
			log.Fatal(err)
			break
		}
		for index := range each_record {
			fmt.Printf("%s ", each_record[index])
		}
		fmt.Println()
	}

}
