package model

type Tracker struct {
	ID        string     `json:"id" mapstructure:"id" bson:"id,omitempty"`
	VehicleID string     `json:"vehicle_id" mapstructure:"vehicle_id" bson:"vehicle_id,omitempty"`
	Location  []Location `json:"location" mapstructure:"location" bson:"location,omitempty"`
	CreatedAt int64      `json:"created_at" mapstructure:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64      `json:"updated_at" mapstructure:"updated_at" bson:"updated_at,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"latitude" mapstructure:"latitude" bson:"latitude,omitempty"`
	Longitude float64 `json:"longitude" mapstructure:"longitude" bson:"longitude,omitempty"`
}
