package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-playground/validator/v10"
)

// Create a new product with validator from domain product
func (pu ProductUsecase) Create(ctx context.Context, product *domain.Product) error {
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		return err
	}
	return pu.productRepo.Create(ctx, product)
}
