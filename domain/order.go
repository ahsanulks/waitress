package domain

import "time"

// OrderState to represent enum of state of the order
type OrderState byte

const (
	// Pending is state that order already created but not paid
	Pending OrderState = iota + 1
	// Paid is state that order already paid but still waiting seller response
	Paid
	// Accepted is state that order already approved by seller
	Accepted
	// Delivered is state that order in delivery courier
	Delivered
	// Remitted is state that order is already done and money already sent to seller
	Remitted
	// Refunded is state that money already sent back to buyer
	Refunded
	// Rejected is state that order got rejected by seller
	Rejected
	// Cancelled is state that order got cancelled by buyer
	Cancelled
	// Expired is state that order not paid after reached the expire time
	Expired
)

type Order struct {
	ID         int         `json:"id"`
	Code       string      `json:"code"`
	BuyerID    int         `json:"buyer_id"`
	SellerID   int         `json:"seller_id"`
	State      OrderState  `json:"state"`
	TotalPrice uint        `json:"total_price"`
	Note       string      `json:"note"`
	Items      []OrderItem `json:"items" autosave:"true"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID         int       `json:"id"`
	OrderID    int       `json:"-"`
	ProductID  int       `json:"-"`
	Product    Product   `json:"product"`
	CartItemID int       `json:"-"`
	Quantity   uint      `json:"quantity"`
	Price      uint      `json:"price"`
	Weight     uint      `json:"weight"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderParams struct {
	CartItemIDs []int  `json:"cart_item_ids" validate:"required,gt=0"`
	BuyerID     int    `json:"buyer_id" validate:"required"`
	Note        string `json:"note"`
}
