package domain

import "time"

type Cart struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CartItems []CartItem `json:"item"`
}

type CartItem struct {
	ID         int       `json:"id"`
	CartID     int       `json:"cart_id"`
	ProductID  int       `json:"product_id"`
	Quantity   uint      `json:"quantity"`
	Weight     uint      `json:"weight"`
	Price      uint      `json:"price"`
	Purchashed bool      `json:"purchased"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
