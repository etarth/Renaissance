package repositories

import "Backend/internal/domain/entities"

type IWishlistRepository interface {
	GetAllWishlistsByUserId(userId string) ([]entities.Wishlist, error)
	InsertNewWishlist(data entities.Wishlist) bool
	// GetArtistById(artistId string) (*entities.Artist, error)
	// GetArtistByUserId(userId string) (*entities.Artist, error)
	// UpdateArtistById(newData entities.Artist, artistId string) error
	// UpdateArtistByUserId(newData entities.Artist, userId string) error
	// DeleteArtistById(artistId string) error
	// DeleteArtistByUserId(userId string) error
}
