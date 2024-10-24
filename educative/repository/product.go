package repository

// Product is an entity, it has attributes of a product
type Product struct {
	id    string
	price float64
	name  string
}

// NewProduct is the only way to create a product object
func NewProduct(id string, price float64, name string) Product {
	return Product{
		id:    id,
		name:  name,
		price: price,
	}
}
