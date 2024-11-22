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
