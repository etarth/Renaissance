package handlers

import (
	"Backend/internal/domain/usecases"
	"Backend/internal/interface/dtos"
	"Backend/pkg/response"
	"Backend/pkg/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ArtworkHandler struct {
	artworkUsecase usecases.IArtworkUsecase
	validator      validator.DTOValidator
}

func NewArtworkHandler(usecases usecases.IArtworkUsecase, validator validator.DTOValidator) *ArtworkHandler {
	return &ArtworkHandler{
		artworkUsecase: usecases,
		validator:      validator,
	}
}

func (h *ArtworkHandler) GetAllArtworks(c *fiber.Ctx) error {
	resReturn, err := h.artworkUsecase.GetAllArtworks()
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	resp := response.NewResponseFactory(response.SUCCESS, resReturn)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtworkHandler) GetArtworkById(c *fiber.Ctx) error {
	artworkId := c.Params("id")
	if artworkId == "" {
		resp := response.NewResponseFactory(response.ERROR, "Artwork ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	req := &dtos.ArtworkDTO{}
	artwork, err := h.artworkUsecase.GetArtworkById(req, artworkId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	resp := response.NewResponseFactory(response.SUCCESS, artwork)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ArtworkHandler) InsertNewArtwork(c *fiber.Ctx) error {
	var createArtworkDTO dtos.InsertNewArtworkDTO
	if err := c.BodyParser(&createArtworkDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(createArtworkDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.artworkUsecase.InsertNewArtwork(&createArtworkDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, createArtworkDTO)
	return resp.SendResponse(c, fiber.StatusCreated)
}
