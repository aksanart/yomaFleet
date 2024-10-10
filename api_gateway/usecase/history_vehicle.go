package usecase

import (
	"context"
	"net/http"

	"github.com/aksan/weplus/apigw/contract"
	uerror "github.com/aksan/weplus/apigw/pkg/error"
	trackerContract "github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	vehicleContract "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
)

// HistoryLocationVehicle implements transport.UseCaseContract.
func (u *UseCase) HistoryLocationVehicle(ctx context.Context, request *contract.HistoryLocationVehicleReq) (response *contract.HistoryLocationVehicleResponse, err error) {
	logData := map[string]any{
		"method":  "HistoryLocationVehicle",
		"request": request,
	}
	_, err = u.validateSession(ctx)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-validateSession", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	if request.GetVehicleId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}

	//get vehicle
	vehicle, err := u.repo.VehicleService.DetailVehicle(ctx, &vehicleContract.IDVehicleReq{
		Id: request.GetVehicleId(),
	})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-DetailVehicle", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	if vehicle.GetData() == nil {
		return &contract.HistoryLocationVehicleResponse{
			Code:    http.StatusOK,
			Message: "success",
			Data:    nil,
		}, nil
	}
	tracker, err := u.repo.TrackerService.DetailTracker(ctx, &trackerContract.IDTrackerReq{Id: vehicle.GetData().GetId()})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-DetailTracker", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	var locations []*contract.Location
	for _, v := range tracker.GetData().GetLocation() {
		location := contract.Location{
			Latitude:  v.GetLatitude(),
			Longitude: v.GetLongitude(),
		}
		locations = append(locations, &location)
	}

	logger.GetLogger().Debug("success", logger.ConvertMapToFields(logData)...)
	return &contract.HistoryLocationVehicleResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data: &contract.HistoryLocationVehicleResponse_Data{
			VehicleId:     vehicle.Data.GetId(),
			VehicleName:   vehicle.Data.GetVehicleName(),
			VehicleModel:  vehicle.Data.GetVehicleModel(),
			VehicleStatus: vehicle.Data.GetVehicleStatus(),
			Mileage:       vehicle.Data.GetMileage(),
			LicenseNumber: vehicle.Data.GetLicenseNumber(),
			Location:      locations,
		},
	}, nil
}
