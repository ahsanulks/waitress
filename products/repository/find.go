package repository

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

type findProduct struct {
	db rel.Repository
}

// FindAll product with offset and limit
func (pr ProductRepository) FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product
	err := pr.db.FindAll(ctx, &products, rel.Offset(offset), rel.Limit(limit))
	return products, err
}
