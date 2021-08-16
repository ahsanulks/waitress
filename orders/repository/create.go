package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
)

// Create will create order
func (or OrderRepository) Create(ctx context.Context, orderPrams domain.OrderParams) (domain.Order, error) {
	var (
		order      domain.Order
		orderItems []domain.OrderItem
		cartItems  []domain.CartItem
		totalPrice uint
	)
	err := or.db.Transaction(ctx, func(ctx context.Context) error {
		or.db.FindAll(ctx, &cartItems, where.InInt("id", orderPrams.CartItemIDs), sort.Asc("product_id"), rel.ForUpdate())

		if len(orderPrams.CartItemIDs) != len(cartItems) {
			return errors.New("cannot found all cart item")
		}

		// create order item and check the stock
		for i, cartItem := range cartItems {
			var product domain.Product
			// lock the product for update
			if err := or.db.Find(ctx, &product, rel.Eq("id", cartItem.ProductID), rel.ForUpdate()); err != nil {
				return errors.New("product not found")
			}
			if cartItem.Quantity > product.Stock {
				return errors.New("stock not enough")
			}
			// decrease the stock
			product.Stock = product.Stock - cartItem.Quantity

			// for update cart item that we literate
			cartItems[i].Product = &product
			cartItems[i].Purchased = true

			if err := or.db.Update(ctx, &cartItems[i]); err != nil {
				return err
			}

			// calculate total price
			totalPrice = totalPrice + (cartItem.Quantity * product.Price)
			orderItems = append(orderItems, buildOrderItem(cartItem, product))
		}
		order = buildOrder(totalPrice, orderPrams, *cartItems[0].Product)
		order.Items = orderItems
		if err := or.db.Insert(ctx, &order); err != nil {
			return err
		}
		// format OS (online store) - 2 last digit year - 10 digit zero + id
		orderCode := fmt.Sprintf("OS%02d%010d", time.Now().Year()-2000, order.ID)

		return or.db.Update(ctx, &order, rel.Set("code", orderCode))
	})
	if err != nil {
		return domain.Order{}, err
	}
	return order, err
}

func buildOrderItem(cartItem domain.CartItem, product domain.Product) domain.OrderItem {
	return domain.OrderItem{
		ProductID:  product.ID,
		CartItemID: cartItem.ID,
		Quantity:   cartItem.Quantity,
		Price:      product.Price,
		Weight:     product.Weight,
		Product:    product,
	}
}

func buildOrder(totalPrice uint, orderPrams domain.OrderParams, product domain.Product) domain.Order {
	return domain.Order{
		BuyerID:    orderPrams.BuyerID,
		SellerID:   product.SellerID,
		State:      domain.Pending,
		TotalPrice: totalPrice,
		Note:       orderPrams.Note,
	}
}
