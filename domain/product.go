package domain

import "time"

// Product struct represent product domain
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	SellerID  int       `json:"seller_id" validate:"required"`
	Price     uint      `json:"price" validate:"required"`
	Active    bool      `json:"active"`
	Stock     uint      `json:"stock" validate:"required"`
	Weight    uint      `json:"weight" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
