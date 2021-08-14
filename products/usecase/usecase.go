package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

type ProductUsecase struct {
	productRepo ProductRepository
}

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error)
}

func NewProductUsecase(productRepo ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo}
}
