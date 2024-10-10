package model

type DtoLiveTrackingResponse struct {
	VehicleId   string `json:"vehicle_id" mapstructure:"vehicle_id"`
	Latitude    string `json:"latitude" mapstructure:"latitude"`
	Longitude   string `json:"longitude" mapstructure:"longitude"`
}
