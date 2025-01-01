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

// func (r *artworkRepository) GetArtworkById(artworkId string) (*entities.Artwork, error) {
// 	var artwork entities.Artwork
// 	filter := bson.M{"artwork_id": artworkId}

// 	err := r.Collection.FindOne(r.Context, filter).Decode(&artwork)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &artwork, nil
// }

func (r *orderItemsRepository) InsertNewOrderItems(data entities.OrderItems) bool {
	if _, err := r.Collection.InsertOne(r.Context, data); err != nil {
		fiberlog.Errorf("OrderItems -> InsertNewOrderItems: %s \n", err)
		return false
	}
	return true
}

// func (r *artworkRepository) UpdateArtworkById(newData entities.Artwork, artworkId string) error {
// 	artworkData := bson.M{"artwork_id": artworkId}

// 	_, err := r.Collection.UpdateOne(r.Context, artworkData, bson.M{"$set": newData})

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}
// 	return nil
// }

// func (r *artistRepository) GetArtistByUserId(userId string) (*entities.Artist, error) {
// 	var artist entities.Artist
// 	filter := bson.M{"user_id": userId}

// 	err := r.Collection.FindOne(r.Context, filter).Decode(&artist)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &artist, nil
// }

// func (r *artistRepository) UpdateArtistByUserId(newData entities.Artist, userId string) error {
// 	artistData := bson.M{"user_id": userId}

// 	_, err := r.Collection.UpdateOne(r.Context, artistData, bson.M{"$set": newData})

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }

// func (r *artworkRepository) DeleteArtworkById(artworkId string) error {
// 	artworkData := bson.M{"artwork_id": artworkId}

// 	_, err := r.Collection.DeleteOne(r.Context, artworkData)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }

// func (r *artistRepository) DeleteArtistByUserId(userId string) error {
// 	artistData := bson.M{"user_id": userId}

// 	_, err := r.Collection.DeleteOne(r.Context, artistData)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }
