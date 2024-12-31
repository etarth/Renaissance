package handlers

import (
	"Backend/internal/domain/usecases"
	"Backend/internal/interface/dtos"
	"Backend/pkg/response"
	"Backend/pkg/validator"

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

func (h *OrderItemsHandler) GetAllOrderItems(c *fiber.Ctx) error {
	resReturn, err := h.orderItemsUsecase.GetAllOrderItems()
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	resp := response.NewResponseFactory(response.SUCCESS, resReturn)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *OrderItemsHandler) GetAllOrderItemsByOrderId(c *fiber.Ctx) error {
	orderId := c.Query("order_id")
	if orderId == "" {
		resp := response.NewResponseFactory(response.ERROR, "User ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	req := &dtos.OrderItemsDTO{}
	orderItems, err := h.orderItemsUsecase.GetAllOrderItemsByOrderId(req, orderId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}
	resp := response.NewResponseFactory(response.SUCCESS, orderItems)
	return resp.SendResponse(c, fiber.StatusOK)
}

// func (h *ArtworkHandler) GetArtworkById(c *fiber.Ctx) error {
// 	artworkId := c.Query("artwork_id")
// 	if artworkId == "" {
// 		resp := response.NewResponseFactory(response.ERROR, "Artwork ID is required")
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}

// 	req := &dtos.ArtworkDTO{}
// 	artwork, err := h.artworkUsecase.GetArtworkById(req, artworkId)
// 	if err != nil {
// 		resp := response.NewResponseFactory(response.ERROR, err.Error())
// 		return resp.SendResponse(c, fiber.StatusInternalServerError)
// 	}

// 	resp := response.NewResponseFactory(response.SUCCESS, artwork)
// 	return resp.SendResponse(c, fiber.StatusOK)
// }

// func (h *ArtworkHandler) InsertNewArtwork(c *fiber.Ctx) error {
// 	var createArtworkDTO dtos.InsertNewArtworkDTO
// 	if err := c.BodyParser(&createArtworkDTO); err != nil {
// 		resp := response.NewResponseFactory(response.ERROR, err.Error())
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}

// 	if errs := h.validator.Validate(createArtworkDTO); len(errs) > 0 {
// 		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}

// 	apperr := h.artworkUsecase.InsertNewArtwork(&createArtworkDTO)
// 	if apperr != nil {
// 		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
// 		return resp.SendResponse(c, apperr.HttpCode)
// 	}

// 	resp := response.NewResponseFactory(response.SUCCESS, createArtworkDTO)
// 	return resp.SendResponse(c, fiber.StatusCreated)
// }

// func (h *ArtworkHandler) UpdateArtworkById(c *fiber.Ctx) error {
// 	artworkId := c.Query("artwork_id")
// 	if artworkId == "" {
// 		resp := response.NewResponseFactory(response.ERROR, "artwork ID is required")
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}

// 	var updateArtworkDTO dtos.UpdateArtworkByIdDTO
// 	if err := c.BodyParser(&updateArtworkDTO); err != nil {
// 		resp := response.NewResponseFactory(response.ERROR, err.Error())
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}

// 	if errs := h.validator.Validate(updateArtworkDTO); len(errs) > 0 {
// 		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}

// 	apperr := h.artworkUsecase.UpdateArtworkById(updateArtworkDTO, artworkId)
// 	if apperr != nil {
// 		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
// 		return resp.SendResponse(c, apperr.HttpCode)
// 	}

// 	resp := response.NewResponseFactory(response.SUCCESS, updateArtworkDTO)
// 	return resp.SendResponse(c, fiber.StatusOK)
// }

// func (h *ArtworkHandler) DeleteArtworkById(c *fiber.Ctx) error {
// 	artworkId := c.Query("artwork_id")
// 	if artworkId == "" {
// 		resp := response.NewResponseFactory(response.ERROR, "artwork ID is required")
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}

// 	apperr := h.artworkUsecase.DeleteArtworkById(artworkId)
// 	if apperr != nil {
// 		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
// 		return resp.SendResponse(c, apperr.HttpCode)
// 	}

// 	resp := response.NewResponseFactory(response.SUCCESS, "Artwork deleted successfully")
// 	return resp.SendResponse(c, fiber.StatusOK)
// }
