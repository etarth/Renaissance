package usecases

import (
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
)

type IOrderItemsUsecase interface {
	GetAllOrderItemsByField(req *dtos.OrderItemsDTO, field string, id string) ([]dtos.OrderItemsDTO, *apperror.AppError)
	GetOrderItemsById(req *dtos.OrderItemsDTO, orderItemsId string) (*dtos.OrderItemsDTO, *apperror.AppError)
	InsertNewOrderItems(dto *dtos.InsertNewOrderItemsDTO) *apperror.AppError
	UpdateOrderItemsById(newData dtos.UpdateOrderItemsByIdDTO, orderItemsId string) *apperror.AppError
	DeleteOrderItemsById(orderItemsId string) *apperror.AppError
}
