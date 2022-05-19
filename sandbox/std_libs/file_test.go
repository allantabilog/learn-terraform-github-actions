package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

type Metadata struct {
	FileName string
	Inode    int
}

type MetadataReader interface {
	GetFileName() string
	GetInode() int
}

func (metadata Metadata) GetFileName() string {
	return metadata.FileName
}

func (metadata Metadata) GetInode() int {
	return metadata.Inode
}

func TestSandbox(t *testing.T) {
	var metadata Metadata
	metadata = Metadata{FileName: "/tmp/test.txt", Inode: 1000}

	//var reader MetadataReader

	fmt.Println(metadata.GetInode())
	fmt.Println(metadata.GetFileName())

	//fmt.Println(reader.(MetadataReader))
}

func TestFileOpen(t *testing.T) {

	file, err := os.Open("/Users/atabilog/dev/golang/diagrams/allocation-logic-nonexistent.puml")
	if err != nil {
		log.Fatal(err)

		// check if err is / contains a *os.PathError

	}

	fmt.Printf("Opened file %s", file.Name())

	data := make([]byte, 1000)
	count, err := file.Read(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes: %q\n", count, data[:count])

}
