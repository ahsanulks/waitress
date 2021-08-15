package repository

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

// Create cart domain
func (cr CartRepository) Create(ctx context.Context, cart *domain.Cart) error {
	return cr.db.Insert(ctx, cart)
}
