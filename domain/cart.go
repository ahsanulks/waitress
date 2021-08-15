package domain

import "time"

type Cart struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CartItems []CartItem `json:"items,omitempty"`
}

type CartItem struct {
	ID        int       `json:"id"`
	CartID    int       `json:"-"`
	ProductID int       `json:"-"`
	Product   *Product  `json:"product,omitempty" autosave:"true"`
	Quantity  uint      `json:"quantity"`
	Purchased bool      `json:"purchased"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartItemParams struct {
	CartID    int  `json:"cart_id" validate:"required"`
	ProductID int  `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"gte=1"`
}
