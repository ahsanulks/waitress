package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

type ProductUsecase struct {
	productRepo ProductRepository
}

// ProductRepository method that need by product usecase.
type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error)
}

// NewProductUsecase create new product usecase.
func NewProductUsecase(productRepo ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo}
}
