package repository

// ProductRepository defines all functionalities in a repository
type ProductRepository interface {

	// GetAllProducts allows to search for all products
	GetAllProducts() []Product
}
