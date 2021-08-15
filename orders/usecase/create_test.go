package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ahsanulks/waitress/domain"
	"github.com/ahsanulks/waitress/orders/usecase"
)

func TestOrderUsecase_Create(t *testing.T) {
	tests := []struct {
		name       string
		orderPrams domain.OrderParams
		want       domain.Order
		wantErr    bool
	}{
		{
			name: "success create",
			orderPrams: domain.OrderParams{
				BuyerID:     2,
				CartItemIDs: []int{1},
			},
			want: domain.Order{
				ID:    2,
				State: domain.Pending,
			},
			wantErr: false,
		},
		{
			name: "error when create",
			orderPrams: domain.OrderParams{
				BuyerID:     3,
				CartItemIDs: []int{1},
			},
			want:    domain.Order{},
			wantErr: true,
		},
		{
			name: "error when validate data",
			orderPrams: domain.OrderParams{
				CartItemIDs: []int{1},
			},
			want:    domain.Order{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ou  = usecase.NewOrderUsecase(fakeOrderRepository{})
				ctx = context.TODO()
			)
			got, err := ou.Create(ctx, tt.orderPrams)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
