package repositories

import "Backend/internal/domain/entities"

type IArtworkRepository interface {
	GetAllArtworks() ([]entities.Artwork, error)
	GetArtworkById(artworkId string) (*entities.Artwork, error)
	// GetArtistByUserId(userId string) (*entities.Artist, error)
	InsertNewArtwork(data entities.Artwork) bool
	UpdateArtworkById(newData entities.Artwork, artworkId string) error
	// UpdateArtistByUserId(newData entities.Artist, userId string) error
	DeleteArtworkById(artworkId string) error
	// DeleteArtistByUserId(userId string) error
}
