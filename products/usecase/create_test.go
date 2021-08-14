package usecase_test

import (
	"context"
	"testing"

	"github.com/ahsanulks/waitress/domain"
	"github.com/ahsanulks/waitress/products/usecase"
)

func TestProductUsecase_Create(t *testing.T) {
	tests := []struct {
		name    string
		args    *domain.Product
		wantErr bool
	}{
		{
			name: "success create",
			args: &domain.Product{
				Name:     "test product",
				SellerID: 10,
				Price:    10000,
				Stock:    99,
				Weight:   1000,
			},
			wantErr: false,
		},
		{
			name: "error when create product",
			args: &domain.Product{
				Name:     "test product",
				SellerID: 2,
				Price:    10000,
				Stock:    99,
				Weight:   1000,
			},
			wantErr: true,
		},
		{
			name: "error when validate struct",
			args: &domain.Product{
				SellerID: 2,
				Price:    10000,
				Stock:    99,
				Weight:   1000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				pu  = usecase.NewProductUsecase(fakeProductRepository{})
				ctx = context.TODO()
			)
			if err := pu.Create(ctx, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("ProductUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
