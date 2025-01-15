package handlers

import (
	"Backend/internal/domain/usecases"
	"Backend/internal/interface/dtos"
	"Backend/pkg/response"
	"Backend/pkg/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ArtistHandler struct {
	artistUsecase usecases.IArtistUsecase
	validator     validator.DTOValidator
}

func NewArtistHandler(usecases usecases.IArtistUsecase, validator validator.DTOValidator) *ArtistHandler {
	return &ArtistHandler{
		artistUsecase: usecases,
		validator:     validator,
	}
}

func (h *ArtistHandler) GetAllArtists(c *fiber.Ctx) error {
	resReturn, err := h.artistUsecase.GetAllArtists()
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	resp := response.NewResponseFactory(response.SUCCESS, resReturn)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtistHandler) GetArtistById(c *fiber.Ctx) error {
	artistId := c.Params("artistId")
	resReturn, err := h.artistUsecase.GetArtistById(artistId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	resp := response.NewResponseFactory(response.SUCCESS, resReturn)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtistHandler) GetArtistByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	resReturn, err := h.artistUsecase.GetArtistByUserId(userId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	resp := response.NewResponseFactory(response.SUCCESS, resReturn)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtistHandler) InsertNewArtist(c *fiber.Ctx) error {
	var createArtistDTO dtos.InsertNewArtistDTO
	if err := c.BodyParser(&createArtistDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(createArtistDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.artistUsecase.InsertNewArtist(&createArtistDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, createArtistDTO)
	return resp.SendResponse(c, fiber.StatusCreated)
}

func (h *ArtistHandler) UpdateArtistById(c *fiber.Ctx) error {
	var updateArtistDTO dtos.ArtistDTO
	if err := c.BodyParser(&updateArtistDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(updateArtistDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.artistUsecase.UpdateArtistById(updateArtistDTO, c.Params("artistId"))
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, updateArtistDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtistHandler) UpdateArtistByUserId(c *fiber.Ctx) error {
	var updateArtistDTO dtos.ArtistDTO
	if err := c.BodyParser(&updateArtistDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(updateArtistDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.artistUsecase.UpdateArtistByUserId(updateArtistDTO, c.Params("userId"))
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, updateArtistDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtistHandler) DeleteArtistById(c *fiber.Ctx) error {
	artistId := c.Params("artistId")
	apperr := h.artistUsecase.DeleteArtistById(artistId)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}
	resp := response.NewResponseFactory(response.SUCCESS, nil)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtistHandler) DeleteArtistByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	apperr := h.artistUsecase.DeleteArtistByUserId(userId)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}
	resp := response.NewResponseFactory(response.SUCCESS, nil)
	return resp.SendResponse(c, fiber.StatusOK)
}
