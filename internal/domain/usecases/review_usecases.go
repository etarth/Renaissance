package usecases

import (
	"Backend/internal/domain/entities"
	"Backend/internal/interface/dtos"
	"Backend/internal/interface/repositories"
	"Backend/pkg/apperror"
	"Backend/pkg/config"
	"reflect"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type reviewUsecase struct {
	cfg              config.Config
	logger           *zap.Logger
	reviewRepository repositories.IReviewRepository
}

func NewReviewUsecases(cfg config.Config, logger *zap.Logger, reviewRepository repositories.IReviewRepository) IReviewUsecase {
	return &reviewUsecase{
		cfg:              cfg,
		logger:           logger,
		reviewRepository: reviewRepository,
	}
}

func (u *reviewUsecase) GetAllReviews() ([]dtos.ReviewDTO, *apperror.AppError) {
	reviews, err := u.reviewRepository.GetAllReviews()
	if err != nil {
		return nil, apperror.InternalServerError("review not found")
	}
	if len(reviews) == 0 {
		return []dtos.ReviewDTO{}, nil
	}
	res := make([]dtos.ReviewDTO, len(reviews))
	for i := 0; i < len(reviews); i++ {
		res[i] = dtos.ReviewDTO{
			ReviewId:  (reviews)[i].ReviewId,
			UserId:    (reviews)[i].UserId,
			ArtistId:  (reviews)[i].ArtistId,
			Rating:    (reviews)[i].Rating,
			Comment:   (reviews)[i].Comment,
			CreatedAt: (reviews)[i].CreatedAt,
		}
	}
	return res, nil
}

func (u *reviewUsecase) GetReviewsByArtistId(artistId string) ([]dtos.ReviewDTO, *apperror.AppError) {
	reviews, err := u.reviewRepository.GetReviewsByArtistId(artistId)
	if err != nil {
		return nil, apperror.InternalServerError("failed to retrieve reviews by artistId")
	}

	if len(reviews) == 0 {
		return []dtos.ReviewDTO{}, nil // Return an empty slice if no reviews
	}

	res := make([]dtos.ReviewDTO, len(reviews))
	for i := 0; i < len(reviews); i++ {
		res[i] = dtos.ReviewDTO{
			ReviewId:  reviews[i].ReviewId,
			UserId:    reviews[i].UserId,
			ArtistId:  reviews[i].ArtistId,
			Rating:    reviews[i].Rating,
			Comment:   reviews[i].Comment,
			CreatedAt: reviews[i].CreatedAt,
		}
	}

	return res, nil
}

func (u *reviewUsecase) GetReviewsByUserId(userId string) ([]dtos.ReviewDTO, *apperror.AppError) {
	reviews, err := u.reviewRepository.GetReviewsByUserId(userId)
	if err != nil {
		return nil, apperror.InternalServerError("failed to retrieve reviews by userId")
	}

	if len(reviews) == 0 {
		return []dtos.ReviewDTO{}, nil // Return an empty slice if no reviews
	}

	res := make([]dtos.ReviewDTO, len(reviews))
	for i := 0; i < len(reviews); i++ {
		res[i] = dtos.ReviewDTO{
			ReviewId:  reviews[i].ReviewId,
			UserId:    reviews[i].UserId,
			ArtistId:  reviews[i].ArtistId,
			Rating:    reviews[i].Rating,
			Comment:   reviews[i].Comment,
			CreatedAt: reviews[i].CreatedAt,
		}
	}

	return res, nil
}

func (u *reviewUsecase) GetReviewById(req *dtos.ReviewDTO, reviewId string) (*dtos.ReviewDTO, *apperror.AppError) {
	review, err := u.reviewRepository.GetReviewById(reviewId)
	if err != nil {
		return nil, apperror.InternalServerError("failed to fetch review")
	}

	if review == nil {
		return nil, apperror.NotFoundError("review not found")
	}

	res := &dtos.ReviewDTO{
		ReviewId:  review.ReviewId,
		UserId:    review.UserId,
		ArtistId:  review.ArtistId,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
	}

	return res, nil
}

func (u *reviewUsecase) InsertNewReview(dto *dtos.InsertNewReviewDTO) *apperror.AppError {
	newReviewId := uuid.New().String()

	newReview := entities.Review{
		ReviewId:  newReviewId,
		ArtistId:  dto.ArtistId,
		UserId:    dto.UserId,
		Rating:    dto.Rating,
		Comment:   dto.Comment,
		CreatedAt: dto.CreatedAt,
	}

	if err := u.reviewRepository.InsertNewReview(newReview); err != true {
		u.logger.Named("CreateReview").Error("Failed to insert review", zap.String("reviewID", dto.ReviewId))
		return apperror.InternalServerError("Failed to insert review")
	}

	u.logger.Named("CreateReview").Info("Success: ", zap.String("review_id", newReview.ReviewId))
	return nil
}

func (u *reviewUsecase) UpdateReviewById(newData dtos.UpdateReviewByIdDTO, reviewId string) *apperror.AppError {
	// Fetch the review by ID
	review, err := u.reviewRepository.GetReviewById(reviewId)
	if err != nil {
		return apperror.InternalServerError("failed to fetch review")
	}
	if review == nil {
		return apperror.NotFoundError("review not found")
	}

	// Reflect over newData and review to apply updates
	newDataValue := reflect.ValueOf(newData)
	reviewValue := reflect.ValueOf(review).Elem() // Dereference pointer

	// Loop through each field of the new data
	for i := 0; i < newDataValue.NumField(); i++ {
		field := newDataValue.Type().Field(i) // Get the field definition
		newValue := newDataValue.Field(i)     // Get the value of the field in newData

		// Only update if the field is not the zero value (empty or nil)
		if !newValue.IsZero() {
			reviewField := reviewValue.FieldByName(field.Name)
			if reviewField.IsValid() && reviewField.CanSet() {
				// Handle both pointer and non-pointer types
				if reviewField.Kind() == reflect.Ptr && newValue.Kind() != reflect.Ptr {
					// If the review field is a pointer but new value is not, we create a new pointer
					reviewField.Set(reflect.New(reviewField.Type().Elem()).Elem())
				}
				// Set the new value in the review field
				reviewField.Set(newValue)
			}
		}
	}

	// Update the review in the repository
	if err := u.reviewRepository.UpdateReviewById(*review, reviewId); err != nil {
		u.logger.Named("UpdateReviewById").Error("Failed to update review", zap.String("review_id", reviewId))
		return apperror.InternalServerError("failed to update review")
	}

	// Log success
	u.logger.Named("UpdateReviewById").Info("Success", zap.String("review_id", reviewId))
	return nil
}

func (u *reviewUsecase) DeleteReviewById(reviewId string) *apperror.AppError {
	review, err := u.reviewRepository.GetReviewById(reviewId)
	if err != nil {
		return apperror.InternalServerError("failed to fetch review")
	}
	if review == nil {
		return apperror.NotFoundError("review not found")
	}

	// Delete the review
	err = u.reviewRepository.DeleteReviewById(reviewId)
	if err != nil {
		u.logger.Named("DeleteReviewById").Error("Failed to delete review", zap.String("review_id", reviewId))
		return apperror.InternalServerError("failed to delete review")
	}

	u.logger.Named("DeleteReviewById").Info("Review deleted successfully", zap.String("review_id", reviewId))
	return nil
}

func (u *reviewUsecase) GetAverageRatingByArtistId(artistId string) (float64, *apperror.AppError) {
	if artistId == "" {
		return 0, apperror.BadRequestError("artist ID is required")
	}

	avgRating, err := u.reviewRepository.GetAverageRatingByArtistId(artistId)
	if err != nil {
		return 0, apperror.InternalServerError("failed to fetch average rating")
	}

	if avgRating == 0 {
		u.logger.Named("GetAverageRatingByArtistId").Info("No reviews found for artist", zap.String("artist_id", artistId))
	}

	return avgRating, nil
}
