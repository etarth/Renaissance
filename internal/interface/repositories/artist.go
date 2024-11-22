package repositories

import "Backend/internal/domain/entities"

type IArtistRepository interface {
	GetAllArtists() ([]entities.Artist, error)
	GetArtistById(artistId string) (*entities.Artist, error)
	GetArtistByUserId(userId string) (*entities.Artist, error)
	InsertNewArtist(data entities.Artist) bool
	UpdateArtistById(newData entities.Artist, artistId string) error
	UpdateArtistByUserId(newData entities.Artist, userId string) error
	DeleteArtistById(artistId string) error
	DeleteArtistByUserId(userId string) error
}
