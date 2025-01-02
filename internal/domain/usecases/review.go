package usecases

import (
	"Backend/internal/interface/dtos"
	"Backend/pkg/apperror"
)

type IReviewUsecase interface {
	GetAllReviews() ([]dtos.ReviewDTO, *apperror.AppError)
	InsertNewReview(dto *dtos.InsertNewReviewDTO) *apperror.AppError
	GetReviewsByArtistId(artistId string) ([]dtos.ReviewDTO, *apperror.AppError)
	GetReviewsByUserId(userId string) ([]dtos.ReviewDTO, *apperror.AppError)
	GetReviewById(req *dtos.ReviewDTO, reviewId string) (*dtos.ReviewDTO, *apperror.AppError)
	UpdateReviewById(newData dtos.UpdateReviewByIdDTO, reviewId string) *apperror.AppError
	DeleteReviewById(reviewId string) *apperror.AppError
	GetAverageRatingByArtistId(artistId string) (float64, *apperror.AppError)
}
