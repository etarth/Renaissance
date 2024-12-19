package usecases

import (
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
)

type IArtworkUsecase interface {
	GetAllArtworks() ([]dtos.ArtworkDTO, *apperror.AppError)
	GetArtworkById(req *dtos.ArtworkDTO, artworkId string) (*dtos.ArtworkDTO, *apperror.AppError)
	// GetArtistByUserId(req *dtos.ArtistDTO, userId string) (*dtos.ArtistDTO, *apperror.AppError)
	InsertNewArtwork(dto *dtos.InsertNewArtworkDTO) *apperror.AppError
	UpdateArtworkById(newData dtos.UpdateArtworkByIdDTO, artworkId string) *apperror.AppError
	// UpdateArtistByUserId(newData dtos.ArtistDTO, userId string) *apperror.AppError
	// DeleteArtistById(artistId string) *apperror.AppError
	// DeleteArtistByUserId(userId string) *apperror.AppError
}
