package postgres

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

type ProductRepository struct {
	db rel.Repository
}

// NewProductRepository create new product repository.
func NewProductRepository(db rel.Repository) *ProductRepository {
	return &ProductRepository{db}
}

func (pr ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	return pr.db.Insert(ctx, product)
}
