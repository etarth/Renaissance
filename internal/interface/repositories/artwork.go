package repositories

import "Backend/internal/domain/entities"

type IArtworkRepository interface {
	GetAllArtworks() ([]entities.Artwork, error)
	GetArtworkById(artworkId string) (*entities.Artwork, error)
	// GetArtistByUserId(userId string) (*entities.Artist, error)
	// InsertNewArtist(data entities.Artist) bool
	// UpdateArtistById(newData entities.Artist, artistId string) error
	// UpdateArtistByUserId(newData entities.Artist, userId string) error
	// DeleteArtistById(artistId string) error
	// DeleteArtistByUserId(userId string) error
}
