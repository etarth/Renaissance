package repositories

import "Backend/internal/domain/entities"

type IReviewRepository interface {
	GetAllReviews() ([]entities.Review, error)
	InsertNewReview(data entities.Review) bool
	UpdateReviewById(newData entities.Review, ReviewId string) error
	GetReviewById(ReviewId string) (*entities.Review, error)
	DeleteReviewById(ReviewId string) error
	GetReviewsByArtistId(artistId string) ([]entities.Review, error)
	GetReviewsByUserId(userId string) ([]entities.Review, error)
	GetAverageRatingByArtistId(artistId string) (float64, error)
	// GetArtistByUserId(userId string) (*entities.Artist, error)
	// InsertNewReview(data entities.Review) bool
	// UpdateReviewById(newData entities.Review, ReviewId string) error
	// UpdateArtistByUserId(newData entities.Artist, userId string) error
	// DeleteReviewById(ReviewId string) error
	// DeleteArtistByUserId(userId string) error
}
