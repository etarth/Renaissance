package repositories

import (
	"Backend/internal/domain/entities"
	. "Backend/pkg/database"
	"context"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type orderItemsRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

func NewOrderItemsRepository(db *MongoDB) IOrderItemsRepository {
	return &orderItemsRepository{
		Context:    db.Context,
		Collection: db.Database.Collection("OrderItems"),
	}
}

func (r *orderItemsRepository) GetAllOrderItemsByField(filter bson.M) ([]entities.OrderItems, error) {
	options := options.Find()

	cursor, err := r.Collection.Find(r.Context, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.Context)

	orderItems := make([]entities.OrderItems, 0)
	for cursor.Next(r.Context) {
		var item entities.OrderItems
		if err := cursor.Decode(&item); err != nil {
			continue
		}
		orderItems = append(orderItems, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *orderItemsRepository) GetOrderItemsById(orderItemsId string) (*entities.OrderItems, error) {
	var orderItem entities.OrderItems
	filter := bson.M{"order_item_id": orderItemsId}

	err := r.Collection.FindOne(r.Context, filter).Decode(&orderItem)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &orderItem, nil
}

func (r *orderItemsRepository) InsertNewOrderItems(data entities.OrderItems) bool {
	if _, err := r.Collection.InsertOne(r.Context, data); err != nil {
		fiberlog.Errorf("OrderItems -> InsertNewOrderItems: %s \n", err)
		return false
	}
	return true
}

func (r *orderItemsRepository) UpdateOrderItemsById(updateFields bson.M, orderItemsId string) error {
	filter := bson.M{"order_item_id": orderItemsId}

	_, err := r.Collection.UpdateOne(r.Context, filter, bson.M{"$set": updateFields})
	if err != nil {
		return err
	}
	return nil
}

func (r *orderItemsRepository) DeleteOrderItemsById(orderItemsId string) error {
	orderItemsData := bson.M{"order_item_id": orderItemsId}

	_, err := r.Collection.DeleteOne(r.Context, orderItemsData)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}
