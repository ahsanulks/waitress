package products

import "time"

type Product struct {
	ID        int
	Name      string
	SellerID  int
	Price     uint
	Active    bool
	Stock     uint
	Weight    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
