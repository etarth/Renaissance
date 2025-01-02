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

type ReviewHandler struct {
	reviewUsecase usecases.IReviewUsecase
	validator     validator.DTOValidator
}

func NewReviewHandler(usecases usecases.IReviewUsecase, validator validator.DTOValidator) *ReviewHandler {
	return &ReviewHandler{
		reviewUsecase: usecases,
		validator:     validator,
	}
}

// func (h *ReviewHandler) GetAllReviews(c *fiber.Ctx) error {
// 	resReturn, err := h.reviewUsecase.GetAllReviews()
// 	if err != nil {
// 		resp := response.NewResponseFactory(response.ERROR, err.Error())
// 		return resp.SendResponse(c, fiber.StatusBadRequest)
// 	}
// 	resp := response.NewResponseFactory(response.SUCCESS, resReturn)
// 	return resp.SendResponse(c, fiber.StatusOK)
// }

func (h *ReviewHandler) GetAllReviews(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	artistId := c.Query("artist_id")

	var resReturn interface{}
	var err *apperror.AppError

	// Handle different query parameters
	if userId == "" && artistId == "" {
		resReturn, err = h.reviewUsecase.GetAllReviews()
	} else if userId != "" {
		resReturn, err = h.reviewUsecase.GetReviewsByUserId(userId)
	} else if artistId != "" {
		resReturn, err = h.reviewUsecase.GetReviewsByArtistId(artistId)
	}

	// If there is an error, handle it
	if err != nil {
		// If no reviews are found, we treat it as a successful response with an empty result
		if len(resReturn.([]dtos.ReviewDTO)) == 0 {
			resp := response.NewResponseFactory(response.SUCCESS, resReturn)
			return resp.SendResponse(c, fiber.StatusOK)
		}

		// Otherwise, treat it as an error (e.g., database failure)
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError) // You could also return a 404 if you prefer
	}

	// If no error, return the reviews
	resp := response.NewResponseFactory(response.SUCCESS, resReturn)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ReviewHandler) GetReviewById(c *fiber.Ctx) error {
	reviewId := c.Params("review_id") // Fetch reviewId from the route parameters
	if reviewId == "" {
		resp := response.NewResponseFactory(response.ERROR, "Review ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}
	req := &dtos.ReviewDTO{}
	review, err := h.reviewUsecase.GetReviewById(req, reviewId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}
	if review == nil {
		resp := response.NewResponseFactory(response.ERROR, "Review not found")
		return resp.SendResponse(c, fiber.StatusNotFound)
	}
	resp := response.NewResponseFactory(response.SUCCESS, review)
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ReviewHandler) InsertNewReview(c *fiber.Ctx) error {
	var createReviewDTO dtos.InsertNewReviewDTO
	if err := c.BodyParser(&createReviewDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(createReviewDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.reviewUsecase.InsertNewReview(&createReviewDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, createReviewDTO)
	return resp.SendResponse(c, fiber.StatusCreated)
}

func (h *ReviewHandler) UpdateReviewById(c *fiber.Ctx) error {
	reviewId := c.Params("review_id") // Fetch reviewId from the URL parameters
	if reviewId == "" {
		resp := response.NewResponseFactory(response.ERROR, "Review ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	var updateReviewDTO dtos.UpdateReviewByIdDTO // DTO for updating review
	if err := c.BodyParser(&updateReviewDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error()) // Parse body to DTO, handle errors
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	// Validate the DTO (you can add custom validation if needed)
	if errs := h.validator.Validate(updateReviewDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", ")) // Return validation errors
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	// Call use case to update the review in the database
	apperr := h.reviewUsecase.UpdateReviewById(updateReviewDTO, reviewId)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error()) // Handle errors from use case
		return resp.SendResponse(c, apperr.HttpCode)
	}

	// Return success response
	resp := response.NewResponseFactory(response.SUCCESS, updateReviewDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}
func (h *ReviewHandler) DeleteReviewById(c *fiber.Ctx) error {
	reviewId := c.Params("review_id")
	if reviewId == "" {
		resp := response.NewResponseFactory(response.ERROR, "Review ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.reviewUsecase.DeleteReviewById(reviewId)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, "Review deleted successfully")
	return resp.SendResponse(c, fiber.StatusOK)
}

func (h *ReviewHandler) GetAverageRatingByArtistId(c *fiber.Ctx) error {
	artistId := c.Params("artist_id")
	if artistId == "" {
		resp := response.NewResponseFactory(response.ERROR, "Artist ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	avgRating, err := h.reviewUsecase.GetAverageRatingByArtistId(artistId)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	if avgRating == 0 {
		resp := response.NewResponseFactory(response.ERROR, "No ratings found for this artist")
		return resp.SendResponse(c, fiber.StatusNotFound)
	}

	resp := response.NewResponseFactory(response.SUCCESS, map[string]float64{"average_rating": avgRating})
	return resp.SendResponse(c, fiber.StatusOK)
}
