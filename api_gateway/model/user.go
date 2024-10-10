package model

type User struct {
	ID        string `json:"id" mapstructure:"id" bson:"id,omitempty"`
	Email     string `json:"email" mapstructure:"email" bson:"email,omitempty"`
	Password  string `json:"password" mapstructure:"password" bson:"password,omitempty"`
	CreatedAt int64  `json:"created_at" mapstructure:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" mapstructure:"updated_at" bson:"updated_at,omitempty"`
}
