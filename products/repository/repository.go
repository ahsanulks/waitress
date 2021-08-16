package repository

import "github.com/go-rel/rel"

type ProductRepository struct {
	db rel.Repository
}

// NewProductRepository create new product repository.
func NewProductRepository(db rel.Repository) *ProductRepository {
	return &ProductRepository{db: db}
}
