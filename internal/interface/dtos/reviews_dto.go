package dtos

import (
	"time"
)

type ReviewDTO struct {
	ReviewId  string    `json:"review_id,omitempty" bson:"review_id,omitempty"`
	UserId    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ArtistId  string    `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	Rating    string    `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment   string    `json:"comment,omitempty" bson:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type InsertNewReviewDTO struct {
	ReviewId  string    `json:"review_id,omitempty" bson:"review_id,omitempty"`
	UserId    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ArtistId  string    `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	Rating    string    `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment   string    `json:"comment,omitempty" bson:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type UpdateReviewByIdDTO struct {
	Rating  string `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment string `json:"comment,omitempty" bson:"comment,omitempty"`
}

type DeleteReviewDTO struct {
	ReviewId string `json:"review_id" bson:"review_id"`
}

type GetReviewsByArtistIdResponse struct {
	Reviews       []ReviewDTO `json:"reviews"`
	AverageRating float64     `json:"average_rating"`
}
