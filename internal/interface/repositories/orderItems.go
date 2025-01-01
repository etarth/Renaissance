package repositories

import (
	"Backend/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type IOrderItemsRepository interface {
	GetAllOrderItemsByField(filter bson.M) ([]entities.OrderItems, error)
	// GetArtworkById(artworkId string) (*entities.Artwork, error)
	// GetArtistByUserId(userId string) (*entities.Artist, error)
	InsertNewOrderItems(data entities.OrderItems) bool
	// UpdateArtworkById(newData entities.Artwork, artworkId string) error
	// UpdateArtistByUserId(newData entities.Artist, userId string) error
	// DeleteArtworkById(artworkId string) error
	// DeleteArtistByUserId(userId string) error
}
