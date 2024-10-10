package model

type CreateVehicleReq struct {
	VehicleName   string  `protobuf:"bytes,1,opt,name=vehicle_name,json=vehicleName,proto3" json:"vehicle_name,omitempty"`
	VehicleModel  string  `protobuf:"bytes,2,opt,name=vehicle_model,json=vehicleModel,proto3" json:"vehicle_model,omitempty"`
	VehicleStatus string  `protobuf:"bytes,3,opt,name=vehicle_status,json=vehicleStatus,proto3" json:"vehicle_status,omitempty"`
	Mileage       int32   `protobuf:"varint,4,opt,name=mileage,proto3" json:"mileage,omitempty"`
	LicenseNumber string  `protobuf:"bytes,5,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
	Latitude      float64 `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float64 `protobuf:"fixed64,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
}
