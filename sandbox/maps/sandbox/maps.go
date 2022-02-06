package sandbox

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string] Vertex

func Main(){
	var m = map[string]Vertex {
		"Bell Labs": {40.6, -77.8},
		"Google": {11.7, -66.8},
	}

	fmt.Println(m)
}
func Main_Map1(){
	m = make(map[string] Vertex)

	m["Bell Labs"] = Vertex{40.111, -74.222}

	fmt.Println	(m)
}