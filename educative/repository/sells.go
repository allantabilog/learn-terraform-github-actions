package repository

// Sells is an aggregate root, it has the attributes of Sells
type Sells struct {
	product    Product
	repository ProductRepository
}

// NewSells is the only way to create a sells object
func NewSells(repository ProductRepository) Sells {
	return Sells{
		repository: repository,
	}
}

// ListProducts returns all products that are in a repository
func (c *Sells) ListProducts() []Product {
	return c.repository.GetAllProducts()
}
