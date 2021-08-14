package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
	"github.com/go-playground/validator/v10"
)

func (ps ProductUsecase) Create(ctx context.Context, product *domain.Product) error {
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		return err
	}
	return ps.Wraiter.Create(ctx, product)
}
