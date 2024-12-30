package handlers

import (
	"Backend/internal/domain/usecases"
	"Backend/internal/interface/dtos"
	"Backend/pkg/response"
	"Backend/pkg/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type WishlistHandler struct {
	wishlistUsecase usecases.IWishlistUsecase
	validator       validator.DTOValidator
}

func NewWishlistHandler(usecases usecases.IWishlistUsecase, validator validator.DTOValidator) *WishlistHandler {
	return &WishlistHandler{
		wishlistUsecase: usecases,
		validator:       validator,
	}
}

func (h *WishlistHandler) GetAllWishlistsByUserId(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	if userId == "" {
		resp := response.NewResponseFactory(response.ERROR, "User ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	req := &dtos.WishlistDTO{}
	wishlist, err := h.wishlistUsecase.GetAllWishlistsByUserId(req, userId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}
	resp := response.NewResponseFactory(response.SUCCESS, wishlist)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *WishlistHandler) InsertNewWishlist(c *fiber.Ctx) error {
	var createWishlistDTO dtos.InsertNewWishlistDTO
	if err := c.BodyParser(&createWishlistDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(createWishlistDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.wishlistUsecase.InsertNewWishlist(&createWishlistDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, createWishlistDTO)
	return resp.SendResponse(c, fiber.StatusCreated)
}

func (h *WishlistHandler) GetWishlistById(c *fiber.Ctx) error {
	favoriteId := c.Query("favorite_id")
	if favoriteId == "" {
		resp := response.NewResponseFactory(response.ERROR, "Wishlist ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	req := &dtos.WishlistDTO{}
	wishlist, err := h.wishlistUsecase.GetWishlistById(req, favoriteId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}
	resp := response.NewResponseFactory(response.SUCCESS, wishlist)
	return resp.SendResponse(c, fiber.StatusOK)
}

// func (h *ArtworkHandler) UpdateArtworkById(c *fiber.Ctx) error {
// 	artworkId := c.Params("artwork_id")
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
// 	artworkId := c.Params("artwork_id")
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
