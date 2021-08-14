package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

func (pu ProductUsecase) FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	return pu.productRepo.FindAll(ctx, limit, offset)
}
