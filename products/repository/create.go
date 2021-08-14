package repository

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

type createProduct struct {
	db rel.Repository
}

func (cp createProduct) Create(ctx context.Context, product *domain.Product) error {
	return cp.db.Insert(ctx, product)
}
