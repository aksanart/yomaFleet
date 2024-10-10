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

// CreateVehicle implements transport.UseCaseContract.
func (u *UseCase) CreateVehicle(ctx context.Context, request *contract.CreateVehicleReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "CreateVehicle",
		"request": request,
	}
	_, err = u.validateSession(ctx)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-validateSession", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	if request.GetVehicleName() == "" || request.GetVehicleModel() == "" || request.GetVehicleStatus() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	vehicleRes, err := u.repo.VehicleService.CreateVehicle(ctx, &vehicleContract.CreateVehicleReq{
		VehicleName:   request.GetVehicleName(),
		VehicleModel:  request.GetVehicleModel(),
		VehicleStatus: request.GetVehicleStatus(),
		Mileage:       request.GetMileage(),
		LicenseNumber: request.GetLicenseNumber(),
	})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-VehicleService.CreateVehicle", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	_, err = u.repo.TrackerService.CreateTracker(ctx, &trackerContract.CreateTrackerReq{
		VehicleId: vehicleRes.GetData().GetId(),
		Location:  []*trackerContract.Location{{Latitude: request.GetLatitude(), Longitude: request.GetLongitude()}},
	})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-TrackerService.CreateTracker", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	logger.GetLogger().Debug("success", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
