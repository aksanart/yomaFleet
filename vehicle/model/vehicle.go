package model

type Vehicle struct {
	ID            string `json:"id" mapstructure:"id" bson:"id,omitempty"`
	VehicleName   string `json:"vehicle_name" mapstructure:"vehicle_name" bson:"vehicle_name,omitempty"`
	VehicleModel  string `json:"vehicle_model" mapstructure:"vehicle_model" bson:"vehicle_model,omitempty"`
	VehicleStatus string `json:"vehicle_status" mapstructure:"vehicle_status" bson:"vehicle_status,omitempty"`
	LicenseNumber string `json:"license_number" mapstructure:"license_number" bson:"license_number,omitempty"`
	Mileage       int32  `json:"mileage" mapstructure:"mileage" bson:"mileage,omitempty"`
	CreatedAt     int64  `json:"created_at" mapstructure:"created_at" bson:"created_at,omitempty"`
	UpdatedAt     int64  `json:"updated_at" mapstructure:"updated_at" bson:"updated_at,omitempty"`
}
