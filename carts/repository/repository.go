package repository

import "github.com/go-rel/rel"

type CartRepository struct {
	db rel.Repository
}

// NewCartRepository to create new cart repository.
func NewCartRepository(db rel.Repository) *CartRepository {
	return &CartRepository{db}
}
