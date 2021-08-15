package repository

import (
	"context"
	"errors"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
)

func (cr CartRepository) AddItem(ctx context.Context, cartItem *domain.CartItem) error {
	return cr.db.Transaction(ctx, func(ctx context.Context) error {
		var (
			cart    domain.Cart
			product domain.Product
		)
		// check that cart relation is present
		if err := cr.db.Find(ctx, &cart, rel.Eq("id", cartItem.CartID)); err != nil {
			return err
		}

		// check that product relation is present
		// lock the product
		if err := cr.db.Find(ctx, &product, rel.Eq("id", cartItem.ProductID), rel.ForUpdate()); err != nil {
			return err
		}

		// check that product stock >= quantity on cart item requested
		if cartItem.Quantity > product.Stock {
			return errors.New("stock not enough")
		}

		return cr.db.Insert(ctx, cartItem)
	})
}
