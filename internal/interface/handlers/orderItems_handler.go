package handlers

import (
	"Backend/internal/domain/usecases"
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
	"Backend/pkg/response"
	"Backend/pkg/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type OrderItemsHandler struct {
	orderItemsUsecase usecases.IOrderItemsUsecase
	validator         validator.DTOValidator
}

func NewOrderItemsHandler(usecases usecases.IOrderItemsUsecase, validator validator.DTOValidator) *OrderItemsHandler {
	return &OrderItemsHandler{
		orderItemsUsecase: usecases,
		validator:         validator,
	}
}

func (h *OrderItemsHandler) GetAllOrderItemsByField(c *fiber.Ctx) error {
	orderId := c.Query("order_id")
	artworkId := c.Query("artwork_id")
	req := &dtos.OrderItemsDTO{}
	var orderItems []dtos.OrderItemsDTO
	var err *apperror.AppError

	if orderId != "" {
		orderItems, err = h.orderItemsUsecase.GetAllOrderItemsByField(req, "order_id", orderId)
	} else if artworkId != "" {
		orderItems, err = h.orderItemsUsecase.GetAllOrderItemsByField(req, "artwork_id", artworkId)
	} else {
		orderItems, err = h.orderItemsUsecase.GetAllOrderItemsByField(req, "all", orderId)
	}

	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	resp := response.NewResponseFactory(response.SUCCESS, orderItems)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *OrderItemsHandler) GetOrderItemsById(c *fiber.Ctx) error {
	orderItemsId := c.Query("orderItems_id")
	if orderItemsId == "" {
		resp := response.NewResponseFactory(response.ERROR, "orderItems ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	req := &dtos.OrderItemsDTO{}
	orderItems, err := h.orderItemsUsecase.GetOrderItemsById(req, orderItemsId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	resp := response.NewResponseFactory(response.SUCCESS, orderItems)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *OrderItemsHandler) InsertNewOrderItems(c *fiber.Ctx) error {
	var createOrderItemsDTO dtos.InsertNewOrderItemsDTO
	if err := c.BodyParser(&createOrderItemsDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(createOrderItemsDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.orderItemsUsecase.InsertNewOrderItems(&createOrderItemsDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, createOrderItemsDTO)
	return resp.SendResponse(c, fiber.StatusCreated)
}

func (h *OrderItemsHandler) UpdateOrderItemsById(c *fiber.Ctx) error {
	orderItemsId := c.Query("orderItems_id")
	if orderItemsId == "" {
		resp := response.NewResponseFactory(response.ERROR, "orderItems ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	var updateorderItemsDTO dtos.UpdateOrderItemsByIdDTO
	if err := c.BodyParser(&updateorderItemsDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(updateorderItemsDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.orderItemsUsecase.UpdateOrderItemsById(updateorderItemsDTO, orderItemsId)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, updateorderItemsDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *OrderItemsHandler) DeleteOrderItemsById(c *fiber.Ctx) error {
	orderItemsId := c.Query("orderItems_id")
	if orderItemsId == "" {
		resp := response.NewResponseFactory(response.ERROR, "orderItems ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.orderItemsUsecase.DeleteOrderItemsById(orderItemsId)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, "orderItems deleted successfully")
	return resp.SendResponse(c, fiber.StatusOK)
}
