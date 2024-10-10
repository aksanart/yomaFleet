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

// DeleteVehicle implements transport.UseCaseContract.
func (u *UseCase) DeleteVehicle(ctx context.Context, request *contract.IDVehicleReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "DeleteVehicle",
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
	_, err = u.repo.VehicleService.DeleteVehicle(ctx, &vehicleContract.IDVehicleReq{Id: request.GetId()})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-DeleteVehicle", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	_, err = u.repo.TrackerService.DeleteTracker(ctx, &trackerContract.IDTrackerReq{Id: request.GetId()})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-DeleteTracker", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	logger.GetLogger().Debug("success", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
