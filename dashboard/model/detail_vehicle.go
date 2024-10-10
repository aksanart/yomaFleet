package model

type HistoryLocationVehicleResponse struct {
	Code    int32                                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *HistoryLocationVehicleResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Title   string
}

type HistoryLocationVehicleResponse_Data struct {
	VehicleId     string      `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	VehicleName   string      `protobuf:"bytes,2,opt,name=vehicle_name,json=vehicleName,proto3" json:"vehicle_name,omitempty"`
	VehicleModel  string      `protobuf:"bytes,3,opt,name=vehicle_model,json=vehicleModel,proto3" json:"vehicle_model,omitempty"`
	VehicleStatus string      `protobuf:"bytes,4,opt,name=vehicle_status,json=vehicleStatus,proto3" json:"vehicle_status,omitempty"`
	Mileage       int32       `protobuf:"varint,5,opt,name=mileage,proto3" json:"mileage,omitempty"`
	LicenseNumber string      `protobuf:"bytes,6,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
	Location      []*Location `protobuf:"bytes,7,rep,name=location,proto3" json:"location,omitempty"`
}

type Location struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}
