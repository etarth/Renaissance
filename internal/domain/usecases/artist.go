package usecases

import (
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
)

type IArtistUsecase interface {
	GetAllArtists() ([]dtos.ArtistDTO, *apperror.AppError)
	// GetArtistById(req *dtos.ArtistDTO, artistId string) (*dtos.ArtistDTO, *apperror.AppError)
	// GetArtistByUserId(req *dtos.ArtistDTO, userId string) (*dtos.ArtistDTO, *apperror.AppError)
	InsertNewArtist(dto *dtos.InsertNewArtistDTO) *apperror.AppError
	// UpdateArtistById(newData dtos.ArtistDTO, artistId string) *apperror.AppError
	// UpdateArtistByUserId(newData dtos.ArtistDTO, userId string) *apperror.AppError
	// DeleteArtistById(artistId string) *apperror.AppError
	// DeleteArtistByUserId(userId string) *apperror.AppError
}
