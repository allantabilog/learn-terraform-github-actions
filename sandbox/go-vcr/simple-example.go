package main

import (
	"fmt"
	"github.com/seborama/govcr"
	)


func main(){
	const example1CassetteName = "MyCassette1"

	vcr := govcr.NewVCR(example1CassetteName, nil)
	vcr.Client.Get("http://example.com/foo")
	fmt.Printf("%+v\n", vcr.Stats())

}
