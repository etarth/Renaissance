package repositories

import "Backend/internal/domain/entities"

type IOrderItemsRepository interface {
	GetAllOrderItems() ([]entities.OrderItems, error)
	GetAllOrderItemsByOrderId(orderId string) ([]entities.OrderItems, error)
	// GetArtworkById(artworkId string) (*entities.Artwork, error)
	// GetArtistByUserId(userId string) (*entities.Artist, error)
	// InsertNewArtwork(data entities.Artwork) bool
	// UpdateArtworkById(newData entities.Artwork, artworkId string) error
	// UpdateArtistByUserId(newData entities.Artist, userId string) error
	// DeleteArtworkById(artworkId string) error
	// DeleteArtistByUserId(userId string) error
}
