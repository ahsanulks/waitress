package usecase

import (
	"context"

	"github.com/ahsanulks/waitress/domain"
)

// FindAll product with offset and limit
func (pu ProductUsecase) FindAll(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	return pu.productRepo.FindAll(ctx, limit, offset)
}
