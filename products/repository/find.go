package repository

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

type findProduct struct {
	db rel.Repository
}

func (fp findProduct) FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product
	err := fp.db.FindAll(ctx, &products, rel.Offset(offset), rel.Limit(limit))
	return products, err
}
