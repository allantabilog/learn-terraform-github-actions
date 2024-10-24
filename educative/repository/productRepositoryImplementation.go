package repository

import "strconv"

// productRepository is the implementation of the repository
type productRepositoryImplementation struct {
}

// NewProductRepositoryImplementation is the only way to create this object
func NewProductRepositoryImplementation() ProductRepository {
	return &productRepositoryImplementation{}
}

// GetAllProducts allows to search for all products
func (p productRepositoryImplementation) GetAllProducts() []Product {
	var products []Product
	for i := 0; i < 10; i++ {
		//products = append(products, NewProduct(string(i), 5644.8, "Product # "+string(i)))
		products = append(products, NewProduct(strconv.Itoa(i), 5644.8, "Product # "+strconv.Itoa(i)))
	}
	return products
}
