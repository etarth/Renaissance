package usecases

import (
	"Backend/internal/domain/entities"
	"Backend/internal/interface/dtos"
	"Backend/internal/interface/repositories"
	"Backend/pkg/apperror"
	"Backend/pkg/config"
	"reflect"
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

func (u *orderItemsUsecase) GetOrderItemsById(req *dtos.OrderItemsDTO, orderItemsId string) (*dtos.OrderItemsDTO, *apperror.AppError) {
	orderItem, err := u.orderItemsRepository.GetOrderItemsById(orderItemsId)
	if err != nil {
		return nil, apperror.InternalServerError("failed to fetch OrderItems")
	}

	if orderItem == nil {
		return &dtos.OrderItemsDTO{}, nil
	}

	res := &dtos.OrderItemsDTO{
		OrderItemId: orderItem.OrderItemId,
		OrderId:     orderItem.OrderId,
		ArtworkId:   orderItem.ArtworkId,
		Quantity:    orderItem.Quantity,
		TotalPrice:  orderItem.TotalPrice,
		CreatedAt:   orderItem.CreatedAt,
	}
	return res, nil
}

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

func (u *orderItemsUsecase) UpdateOrderItemsById(newData dtos.UpdateOrderItemsByIdDTO, orderItemsId string) *apperror.AppError {
	// Fetch the existing order item
	orderItem, err := u.orderItemsRepository.GetOrderItemsById(orderItemsId)
	if err != nil {
		return apperror.InternalServerError("failed to fetch orderItems")
	}
	if orderItem == nil {
		return apperror.NotFoundError("orderItems not found")
	}

	// Prepare for reflection-based update
	newDataValue := reflect.ValueOf(newData)
	newDataType := reflect.TypeOf(newData)
	updateFields := bson.M{} // Map to hold update fields

	// Reflect over newData and extract fields to update
	for i := 0; i < newDataValue.NumField(); i++ {
		field := newDataType.Field(i)
		fieldName := field.Tag.Get("bson") // Extract BSON tag
		fieldValue := newDataValue.Field(i)

		if fieldName != "" && !fieldValue.IsZero() {
			// Handle pointer types
			if fieldValue.Kind() == reflect.Ptr {
				if !fieldValue.IsNil() {
					updateFields[fieldName] = fieldValue.Elem().Interface()
				}
			} else {
				// Handle non-pointer values
				updateFields[fieldName] = fieldValue.Interface()
			}
		}
	}

	// Update the database if there are fields to update
	if len(updateFields) > 0 {
		if err := u.orderItemsRepository.UpdateOrderItemsById(updateFields, orderItemsId); err != nil {
			u.logger.Named("UpdateOrderItemsById").Error("Failed to update orderItems", zap.String("orderItems_id", orderItemsId))
			return apperror.InternalServerError("failed to update orderItems")
		}
	}

	u.logger.Named("UpdateOrderItemsById").Info("Success", zap.String("orderItems_id", orderItemsId))
	return nil
}

func (u *orderItemsUsecase) DeleteOrderItemsById(orderItemsId string) *apperror.AppError {
	orderItem, err := u.orderItemsRepository.GetOrderItemsById(orderItemsId)
	if err != nil {
		return apperror.InternalServerError("failed to fetch orderItems")
	}
	if orderItem == nil {
		return apperror.NotFoundError("orderItems not found")
	}

	err = u.orderItemsRepository.DeleteOrderItemsById(orderItemsId)
	if err != nil {
		u.logger.Named("DeleteOrderItemsById").Error("failed to delete orderItems", zap.String("orderItems_id", orderItemsId))
		return apperror.InternalServerError("failed to delete orderItems")
	}

	u.logger.Named("DeleteOrderItemsById").Info("orderItems deleted successfully", zap.String("orderItems_id", orderItemsId))
	return nil
}
