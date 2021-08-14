package repository

import "github.com/go-rel/rel"

type CartRepository struct {
	db rel.Repository
}

func NewCartRepository(db rel.Repository) *CartRepository {
	return &CartRepository{db}
}
