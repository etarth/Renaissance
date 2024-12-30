package usecases

import (
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
)

type IWishlistUsecase interface {
	GetAllWishlistsByUserId(req *dtos.WishlistDTO, userId string) ([]dtos.WishlistDTO, *apperror.AppError)
	InsertNewWishlist(dto *dtos.InsertNewWishlistDTO) *apperror.AppError
	// GetArtworkById(req *dtos.ArtworkDTO, artworkId string) (*dtos.ArtworkDTO, *apperror.AppError)
	// GetArtistByUserId(req *dtos.ArtistDTO, userId string) (*dtos.ArtistDTO, *apperror.AppError)
	// UpdateArtworkById(newData dtos.UpdateArtworkByIdDTO, artworkId string) *apperror.AppError
	// UpdateArtistByUserId(newData dtos.ArtistDTO, userId string) *apperror.AppError
	// DeleteArtworkById(artworkId string) *apperror.AppError
	// DeleteArtistByUserId(userId string) *apperror.AppError
}
