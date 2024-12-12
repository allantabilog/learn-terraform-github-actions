package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main()  {
	data := []float64{1.0, 1.0,1.0,2.0,1.0,80,2.0}

	median, _ := stats.Median(data)
	mean, _ := stats.Mean(data)

	fmt.Println(median)
	fmt.Println(mean)
}