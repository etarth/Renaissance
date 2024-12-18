package entities

import (
	"time"
)

type Artist struct {
	ArtistId        string `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	UserId          string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Bio             string `json:"bio,omitempty" bson:"bio,omitempty"`
	Website         string `json:"website" bson:"website"`
	SocialLinks     string `json:"social_links" bson:"social_links"`
	ProfileImageURL string `json:"profile_image_url" bson:"profile_image_url"`
}

type Artwork struct {
	ArtworkId   string     `json:"artwork_id,omitempty" bson:"artwork_id,omitempty"`
	ArtistId    string     `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	Title       string     `json:"title,omitempty" bson:"title,omitempty"`
	Description string     `json:"description,omitempty" bson:"description,omitempty"`
	Category    []Category `json:"category,omitempty" bson:"category,omitempty"`
	Style       string     `json:"style,omitempty" bson:"style,omitempty"`
	Width       float32    `json:"width,omitempty" bson:"width,omitempty"`
	Height      float32    `json:"height,omitempty" bson:"height,omitempty"`
	Price       float32    `json:"price,omitempty" bson:"price,omitempty"`
	ImageURL    string     `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Stock       int        `json:"stock,omitempty" bson:"stock,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Category struct {
	CategoryId   string `json:"category_id,omitempty" bson:"category_id,omitempty"`
	CategoryName string `json:"category_name,omitempty" bson:"category_name,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty"`
}
