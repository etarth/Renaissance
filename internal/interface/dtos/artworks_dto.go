package dtos

import (
	"Backend/internal/domain/entities"
	"time"
)

type ArtworkDTO struct {
	ArtworkId   string              `json:"artwork_id,omitempty" bson:"artwork_id,omitempty"`
	ArtistId    string              `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	Title       string              `json:"title,omitempty" bson:"title,omitempty"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	Category    []entities.Category `json:"category,omitempty" bson:"category,omitempty"`
	Style       string              `json:"style,omitempty" bson:"style,omitempty"`
	Width       float32             `json:"width,omitempty" bson:"width,omitempty"`
	Height      float32             `json:"height,omitempty" bson:"height,omitempty"`
	Price       float32             `json:"price,omitempty" bson:"price,omitempty"`
	ImageURL    string              `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Stock       int                 `json:"stock,omitempty" bson:"stock,omitempty"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}

type InsertNewArtworkDTO struct {
	ArtworkId   string              `json:"artwork_id,omitempty" bson:"artwork_id,omitempty"`
	ArtistId    string              `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	Title       string              `json:"title,omitempty" bson:"title,omitempty"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	Category    []entities.Category `json:"category,omitempty" bson:"category,omitempty"`
	Style       string              `json:"style,omitempty" bson:"style,omitempty"`
	Width       float32             `json:"width,omitempty" bson:"width,omitempty"`
	Height      float32             `json:"height,omitempty" bson:"height,omitempty"`
	Price       float32             `json:"price,omitempty" bson:"price,omitempty"`
	ImageURL    string              `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Stock       int                 `json:"stock,omitempty" bson:"stock,omitempty"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}

type UpdateArtworkByIdDTO struct {
	Title       string              `json:"title,omitempty" bson:"title,omitempty"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	Category    []entities.Category `json:"category,omitempty" bson:"category,omitempty"`
	Style       string              `json:"style,omitempty" bson:"style,omitempty"`
	Width       float32             `json:"width,omitempty" bson:"width,omitempty"`
	Height      float32             `json:"height,omitempty" bson:"height,omitempty"`
	Price       float32             `json:"price,omitempty" bson:"price,omitempty"`
	ImageURL    string              `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Stock       int                 `json:"stock,omitempty" bson:"stock,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}
