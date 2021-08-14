package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

type ProductUsecase struct {
	Wraiter ProductWriter
}

type ProductWriter interface {
	Create(ctx context.Context, product *domain.Product) error
}

func NewProductUsecase(writer ProductWriter) *ProductUsecase {
	return &ProductUsecase{writer}
}
