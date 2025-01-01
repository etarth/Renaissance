package usecases

import (
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
)

type IOrderItemsUsecase interface {
	GetAllOrderItemsByField(req *dtos.OrderItemsDTO, field string, id string) ([]dtos.OrderItemsDTO, *apperror.AppError)
	// GetArtworkById(req *dtos.ArtworkDTO, artworkId string) (*dtos.ArtworkDTO, *apperror.AppError)
	// GetArtistByUserId(req *dtos.ArtistDTO, userId string) (*dtos.ArtistDTO, *apperror.AppError)
	InsertNewOrderItems(dto *dtos.InsertNewOrderItemsDTO) *apperror.AppError
	// UpdateArtworkById(newData dtos.UpdateArtworkByIdDTO, artworkId string) *apperror.AppError
	// UpdateArtistByUserId(newData dtos.ArtistDTO, userId string) *apperror.AppError
	// DeleteArtworkById(artworkId string) *apperror.AppError
	// DeleteArtistByUserId(userId string) *apperror.AppError
}
