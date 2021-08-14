package repository_test

import (
	"context"
	"testing"

	"github.com/ahsanulks/waitress/domain"
	"github.com/ahsanulks/waitress/products/repository"
	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
)

func Test_createProduct_Create(t *testing.T) {
	var (
		db      = reltest.New()
		ctx     = context.TODO()
		product = domain.Product{SellerID: 2}
		repo    = repository.NewProductRepository(db)
	)

	db.ExpectInsert().For(&product)
	assert.Nil(t, repo.Create(ctx, &product))
	assert.NotEmpty(t, product.ID)
}
