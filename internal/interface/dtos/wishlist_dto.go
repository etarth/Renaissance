package dtos

import "time"

type WishlistDTO struct {
	FavoriteId string    `json:"favorite_id,omitempty" bson:"favorite_id,omitempty"`
	UserId     string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ArtworkId  string    `json:"artwork_id,omitempty" bson:"artwork_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type InsertNewWishlistDTO struct {
	FavoriteId string    `json:"favorite_id,omitempty" bson:"favorite_id,omitempty"`
	UserId     string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ArtworkId  string    `json:"artwork_id,omitempty" bson:"artwork_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// type UpdateArtistByIdDTO struct {
// 	Bio             string `json:"bio,omitempty" bson:"bio,omitempty"`
// 	Website         string `json:"website" bson:"website"`
// 	SocialLinks     string `json:"social_links" bson:"social_links"`
// 	ProfileImageURL string `json:"profile_image_url" bson:"profile_image_url"`
// }
