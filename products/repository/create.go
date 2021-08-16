package repository

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

// Create a new product.
func (pr ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	return pr.db.Insert(ctx, product)
}
