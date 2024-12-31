package dtos

import (
	"time"
)

// DTO for retrieving order item details
type OrderItemsDTO struct {
	OrderItemId string    `json:"order_item_id,omitempty" bson:"order_item_id,omitempty"`
	OrderId     string    `json:"order_id,omitempty" bson:"order_id,omitempty"`
	ArtworkId   string    `json:"artwork_id,omitempty" bson:"artwork_id,omitempty"`
	Quantity    int       `json:"quantity,omitempty" bson:"quantity,omitempty"`
	TotalPrice  float32   `json:"total_price,omitempty" bson:"total_price,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// DTO for creating new order items
type InsertNewOrderItemsDTO struct {
	OrderId    string  `json:"order_id,omitempty" bson:"order_id,omitempty"`
	ArtworkId  string  `json:"artwork_id,omitempty" bson:"artwork_id,omitempty"`
	Quantity   int     `json:"quantity,omitempty" bson:"quantity,omitempty"`
	TotalPrice float32 `json:"total_price,omitempty" bson:"total_price,omitempty"`
}

// DTO for updating existing order items
type UpdateOrderItemsByIdDTO struct {
	Quantity   *int     `json:"quantity,omitempty" bson:"quantity,omitempty"` // Use pointers for optional fields
	TotalPrice *float32 `json:"total_price,omitempty" bson:"total_price,omitempty"`
}
