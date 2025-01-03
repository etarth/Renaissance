package repositories

import (
	"Backend/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type IOrderItemsRepository interface {
	GetAllOrderItemsByField(filter bson.M) ([]entities.OrderItems, error)
	GetOrderItemsById(orderItemsId string) (*entities.OrderItems, error)
	InsertNewOrderItems(data entities.OrderItems) bool
	UpdateOrderItemsById(updateFields bson.M, orderItemsId string) error
	DeleteOrderItemsById(orderItemsId string) error
}
