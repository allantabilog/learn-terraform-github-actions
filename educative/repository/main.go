package repository

import "fmt"

// main is the start point
func RepositoryMain() {

	sells := NewSells(NewProductRepositoryImplementation())
	products := sells.ListProducts()

	fmt.Println("List of products:")

	for _, v := range products {
		fmt.Printf("%s - %s - %f\n \n", v.id, v.name, v.price)
	}
}
