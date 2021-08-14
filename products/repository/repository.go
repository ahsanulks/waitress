package repository

import "github.com/go-rel/rel"

type ProductRepository struct {
	createProduct
	findProduct
}

// NewProductRepository create new product repository.
func NewProductRepository(db rel.Repository) *ProductRepository {
	return &ProductRepository{
		createProduct: createProduct{db},
		findProduct:   findProduct{db},
	}
}
