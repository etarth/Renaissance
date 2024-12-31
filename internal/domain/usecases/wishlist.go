package usecases

import (
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
)

type IWishlistUsecase interface {
	GetAllWishlistsByUserId(req *dtos.WishlistDTO, userId string) ([]dtos.WishlistDTO, *apperror.AppError)
	InsertNewWishlist(dto *dtos.InsertNewWishlistDTO) *apperror.AppError
	GetWishlistById(req *dtos.WishlistDTO, favoriteId string) (*dtos.WishlistDTO, *apperror.AppError)
	// GetArtistByUserId(req *dtos.ArtistDTO, userId string) (*dtos.ArtistDTO, *apperror.AppError)
	// UpdateArtworkById(newData dtos.UpdateArtworkByIdDTO, artworkId string) *apperror.AppError
	// UpdateArtistByUserId(newData dtos.ArtistDTO, userId string) *apperror.AppError
	DeleteWishlistById(favoriteId string) *apperror.AppError
	// DeleteArtistByUserId(userId string) *apperror.AppError
}
