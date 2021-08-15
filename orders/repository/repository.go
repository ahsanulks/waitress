package repository

import "github.com/go-rel/rel"

type OrderRepository struct {
	db rel.Repository
}

func NewOrderRepository(db rel.Repository) *OrderRepository {
	return &OrderRepository{db: db}
}
