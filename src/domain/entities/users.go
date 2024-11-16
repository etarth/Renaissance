package entities

type UserDataFormat struct {
	UserID          string   `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Username        string   `json:"username,omitempty" bson:"username,omitempty"`
	Email           string   `json:"email,omitempty" bson:"email,omitempty"`
	FullName        string   `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Phone           string   `json:"phone,omitempty" bson:"phone,omitempty"`
	Address         []string `json:"address,omitempty" bson:"address,omitempty"`
	UserType        string   `json:"user_type,omitempty" bson:"user_type,omitempty"`
	ProfileImageURL string   `json:"profile_image_url,omitempty" bson:"profile_image_url,omitempty"`
}
