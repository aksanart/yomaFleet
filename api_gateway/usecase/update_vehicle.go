package usecase

import (
	"context"
	"net/http"

	"github.com/aksan/weplus/apigw/contract"
	uerror "github.com/aksan/weplus/apigw/pkg/error"
	vehicleContract "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
)

// UpdateVehicle implements transport.UseCaseContract.
func (u *UseCase) UpdateVehicle(ctx context.Context, request *contract.UpdateVehicleReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "UpdateVehicle",
		"request": request,
	}
	_, err = u.validateSession(ctx)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-validateSession", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	if request.GetId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	_, err = u.repo.VehicleService.UpdateVehicle(ctx, &vehicleContract.UpdateVehicleReq{
		Id:            request.GetId(),
		VehicleName:   request.GetVehicleName(),
		VehicleModel:  request.GetVehicleModel(),
		VehicleStatus: request.GetVehicleStatus(),
		Mileage:       request.GetMileage(),
		LicenseNumber: request.GetLicenseNumber(),
	})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-UpdateVehicle", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	logger.GetLogger().Debug("success", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
