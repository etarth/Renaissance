package entities

type Artist struct {
	ArtistId        string `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	UserId          string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Bio             string `json:"bio,omitempty" bson:"bio,omitempty"`
	Website         string `json:"website" bson:"website"`
	SocialLinks     string `json:"social_links" bson:"social_links"`
	ProfileImageURL string `json:"profile_image_url" bson:"profile_image_url"`
}
