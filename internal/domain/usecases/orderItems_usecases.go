package usecases

import (
	"Backend/internal/domain/entities"
	"Backend/internal/interface/dtos"
	"Backend/internal/interface/repositories"
	"Backend/pkg/apperror"
	"Backend/pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type orderItemsUsecase struct {
	cfg                  config.Config
	logger               *zap.Logger
	orderItemsRepository repositories.IOrderItemsRepository
}

func NewOrderItemsUsecases(cfg config.Config, logger *zap.Logger, orderItemsRepository repositories.IOrderItemsRepository) IOrderItemsUsecase {
	return &orderItemsUsecase{
		cfg:                  cfg,
		logger:               logger,
		orderItemsRepository: orderItemsRepository,
	}
}

func (u *orderItemsUsecase) GetAllOrderItemsByField(req *dtos.OrderItemsDTO, field string, id string) ([]dtos.OrderItemsDTO, *apperror.AppError) {
	filter := bson.M{field: id}
	if field == "all" {
		filter = bson.M{}
	}
	orderItems, err := u.orderItemsRepository.GetAllOrderItemsByField(filter)
	if err != nil {
		return nil, apperror.InternalServerError("could not retrieve order items")
	}

	res := make([]dtos.OrderItemsDTO, len(orderItems))
	for i, item := range orderItems {
		res[i] = dtos.OrderItemsDTO{
			OrderItemId: item.OrderItemId,
			OrderId:     item.OrderId,
			ArtworkId:   item.ArtworkId,
			Quantity:    item.Quantity,
			TotalPrice:  item.TotalPrice,
			CreatedAt:   item.CreatedAt,
		}
	}
	return res, nil
}

// func (u *artworkUsecase) GetArtworkById(req *dtos.ArtworkDTO, artworkId string) (*dtos.ArtworkDTO, *apperror.AppError) {
// 	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
// 	if err != nil {
// 		return nil, apperror.InternalServerError("failed to fetch artwork")
// 	}

// 	if artwork == nil {
// 		return &dtos.ArtworkDTO{}, nil
// 	}

// 	res := &dtos.ArtworkDTO{
// 		ArtworkId:   artwork.ArtworkId,
// 		ArtistId:    artwork.ArtistId,
// 		Title:       artwork.Title,
// 		Description: artwork.Description,
// 		Category:    artwork.Category,
// 		Style:       artwork.Style,
// 		Width:       artwork.Width,
// 		Height:      artwork.Height,
// 		Price:       artwork.Price,
// 		ImageURL:    artwork.ImageURL,
// 		Stock:       artwork.Stock,
// 		CreatedAt:   artwork.CreatedAt,
// 		UpdatedAt:   artwork.UpdatedAt,
// 	}
// 	return res, nil
// }

func (u *orderItemsUsecase) InsertNewOrderItems(dto *dtos.InsertNewOrderItemsDTO) *apperror.AppError {
	newOrderItems := entities.OrderItems{
		OrderItemId: dto.OrderItemId,
		OrderId:     dto.OrderId,
		ArtworkId:   dto.ArtworkId,
		Quantity:    dto.Quantity,
		TotalPrice:  dto.TotalPrice,
		CreatedAt:   time.Now(),
	}

	if err := u.orderItemsRepository.InsertNewOrderItems(newOrderItems); err != true {
		u.logger.Named("CreateOrderItems").Error("Failed to insert order items", zap.String("orderItemsId", dto.OrderItemId))
		return apperror.InternalServerError("Failed to insert order items")
	}

	u.logger.Named("CreateOrderItems").Info("Success: ", zap.String("orderItemsId", newOrderItems.OrderItemId))
	return nil
}

// func (u *artworkUsecase) UpdateArtworkById(newData dtos.UpdateArtworkByIdDTO, artworkId string) *apperror.AppError {
// 	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
// 	if err != nil {
// 		return apperror.InternalServerError("failed to fetch artwork")
// 	}
// 	if artwork == nil {
// 		return apperror.NotFoundError("artwork not found")
// 	}

// 	// Reflect over newData and artwork to apply updates
// 	newDataValue := reflect.ValueOf(newData)
// 	artworkValue := reflect.ValueOf(artwork).Elem() // Dereference pointer

// 	for i := 0; i < newDataValue.NumField(); i++ {
// 		field := newDataValue.Type().Field(i) // Get the field definition
// 		newValue := newDataValue.Field(i)     // Get the value of the field in newData

// 		// Only update if the field is not the zero value (empty or nil)
// 		if !newValue.IsZero() {
// 			artworkField := artworkValue.FieldByName(field.Name)
// 			if artworkField.IsValid() && artworkField.CanSet() {
// 				// Handle both pointer and non-pointer types
// 				if artworkField.Kind() == reflect.Ptr && newValue.Kind() != reflect.Ptr {
// 					// If artwork field is a pointer but new value is not a pointer, we create a new pointer for the field
// 					artworkField.Set(reflect.New(artworkField.Type().Elem()).Elem())
// 				}
// 				// Update the field with the new value
// 				artworkField.Set(newValue)
// 			}
// 		}
// 	}
// 	artwork.UpdatedAt = time.Now() // Set the updatedAt timestamp

// 	if err := u.artworkRepository.UpdateArtworkById(*artwork, artworkId); err != nil {
// 		u.logger.Named("UpdateArtworkById").Error("Failed to update artwork", zap.String("artwork_id", artworkId))
// 		return apperror.InternalServerError("failed to update artwork")
// 	}

// 	u.logger.Named("UpdateArtworkById").Info("Success", zap.String("artwork_id", artworkId))
// 	return nil
// }

// func (u *artworkUsecase) DeleteArtworkById(artworkId string) *apperror.AppError {
// 	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
// 	if err != nil {
// 		return apperror.InternalServerError("failed to fetch artwork")
// 	}
// 	if artwork == nil {
// 		return apperror.NotFoundError("artwork not found")
// 	}

// 	err = u.artworkRepository.DeleteArtworkById(artworkId)
// 	if err != nil {
// 		u.logger.Named("DeleteArtworkById").Error("Failed to delete artwork", zap.String("artwork_id", artworkId))
// 		return apperror.InternalServerError("failed to delete artwork")
// 	}

// 	u.logger.Named("DeleteArtworkById").Info("Artwork deleted successfully", zap.String("artwork_id", artworkId))
// 	return nil
// }
