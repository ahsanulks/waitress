package repository_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ahsanulks/waitress/domain"
	"github.com/ahsanulks/waitress/orders/repository"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository_Create(t *testing.T) {
	var (
		db        = reltest.New()
		cartItems = []domain.CartItem{
			{
				ID:        2,
				ProductID: 2,
				Quantity:  10,
				Purchased: false,
			},
		}
		productOutOfStock = domain.Product{
			ID:     2,
			Stock:  5,
			Price:  1000,
			Weight: 2000,
		}
		product = domain.Product{
			ID:       2,
			Stock:    100,
			Price:    1000,
			Weight:   2000,
			SellerID: 2,
		}
		order = domain.Order{
			ID: 1,
		}
		resultOrder = domain.Order{
			ID:         1,
			Code:       "OS210000000001",
			BuyerID:    1,
			SellerID:   2,
			State:      domain.Pending,
			TotalPrice: 10000,
		}
	)
	tests := []struct {
		name       string
		orderPrams domain.OrderParams
		want       domain.Order
		wantErr    bool
		mockFunc   func()
	}{
		{
			name: "cannot found all cart_item_ids",
			orderPrams: domain.OrderParams{
				CartItemIDs: []int{1, 2},
			},
			want:    domain.Order{},
			wantErr: true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					db.ExpectFindAll(where.InInt("id", []int{1, 2}), rel.ForUpdate()).Result(cartItems)
				})
			},
		},
		{
			name: "stock not enough",
			orderPrams: domain.OrderParams{
				CartItemIDs: []int{1},
			},
			want:    domain.Order{},
			wantErr: true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					db.ExpectFindAll(where.InInt("id", []int{1}), rel.ForUpdate()).Result(cartItems)
					db.ExpectFind(rel.Eq("id", cartItems[0].ProductID), rel.ForUpdate()).Result(productOutOfStock)
				})
			},
		},
		{
			name: "error when update cart_item",
			orderPrams: domain.OrderParams{
				CartItemIDs: []int{1},
			},
			want:    domain.Order{},
			wantErr: true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					db.ExpectFindAll(where.InInt("id", []int{1}), rel.ForUpdate()).Result(cartItems)
					db.ExpectFind(rel.Eq("id", cartItems[0].ProductID), rel.ForUpdate()).Result(product)
					db.ExpectUpdate().For(&cartItems[0]).Error(errors.New("error"))
				})
			},
		},
		{
			name: "error when insert order",
			orderPrams: domain.OrderParams{
				CartItemIDs: []int{1},
			},
			want:    domain.Order{},
			wantErr: true,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					db.ExpectFindAll(where.InInt("id", []int{1}), rel.ForUpdate()).Result(cartItems)
					db.ExpectFind(rel.Eq("id", cartItems[0].ProductID), rel.ForUpdate()).Result(product)
					db.ExpectUpdate().For(&cartItems[0])
					db.ExpectInsert().ForType("domain.Order").Error(errors.New("error"))
				})
			},
		},
		{
			name: "success insert order",
			orderPrams: domain.OrderParams{
				CartItemIDs: []int{1},
				BuyerID:     1,
			},
			want:    resultOrder,
			wantErr: false,
			mockFunc: func() {
				db.ExpectTransaction(func(repo *reltest.Repository) {
					db.ExpectFindAll(where.InInt("id", []int{1}), rel.ForUpdate()).Result(cartItems)
					db.ExpectFind(rel.Eq("id", cartItems[0].ProductID), rel.ForUpdate()).Result(product)
					db.ExpectUpdate().For(&cartItems[0])
					db.ExpectInsert().ForType("domain.Order")
					db.ExpectUpdate(rel.Set("code", fmt.Sprintf("OS%02d%010d", time.Now().Year()-2000, order.ID)))
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				or  = repository.NewOrderRepository(db)
				ctx = context.TODO()
			)
			tt.mockFunc()
			got, err := or.Create(ctx, tt.orderPrams)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.Code, got.Code)
			assert.Equal(t, tt.want.TotalPrice, got.TotalPrice)
			assert.Equal(t, tt.want.BuyerID, got.BuyerID)
			assert.Equal(t, tt.want.SellerID, got.SellerID)
			assert.Equal(t, tt.want.State, got.State)
		})
	}
}
