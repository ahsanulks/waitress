package repository_test

import (
	"context"
	"testing"

	"github.com/ahsanulks/waitress/carts/repository"
	"github.com/ahsanulks/waitress/domain"
	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestCartRepository_Create(t *testing.T) {
	var (
		db   = reltest.New()
		repo = repository.NewCartRepository(db)
		cart = domain.Cart{UserID: 2}
		ctx  = context.TODO()
	)

	db.ExpectInsert().For(&cart)
	assert.Nil(t, repo.Create(ctx, &cart))
	assert.NotEmpty(t, cart.ID)
}
