package model

type ListVehicleResponse struct {
	Code    int32                       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*ListVehicleResponse_Data `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Title string
}

type ListVehicleResponse_Data struct {
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VehicleName   string `protobuf:"bytes,2,opt,name=vehicle_name,json=vehicleName,proto3" json:"vehicle_name,omitempty"`
	VehicleModel  string `protobuf:"bytes,3,opt,name=vehicle_model,json=vehicleModel,proto3" json:"vehicle_model,omitempty"`
	VehicleStatus string `protobuf:"bytes,4,opt,name=vehicle_status,json=vehicleStatus,proto3" json:"vehicle_status,omitempty"`
	Mileage       int32  `protobuf:"varint,5,opt,name=mileage,proto3" json:"mileage,omitempty"`
	LicenseNumber string `protobuf:"bytes,6,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
}
