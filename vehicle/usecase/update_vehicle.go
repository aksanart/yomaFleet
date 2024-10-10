package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/vehicle/contract"
	"github.com/aksanart/vehicle/model"
	uerror "github.com/aksanart/vehicle/pkg/error"
	"github.com/aksanart/vehicle/pkg/logger"
)

// UpdateVehicle implements transport.UseCaseContract.
func (u *UseCase) UpdateVehicle(ctx context.Context, request *contract.UpdateVehicleReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "UpdateVehicle",
		"request": request,
	}
	if request.GetId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	err = u.repo.MongoDb.Vehicle.Update(
		ctx,
		request.GetId(),
		&model.Vehicle{
			VehicleName:   request.GetVehicleName(),
			VehicleModel:  request.GetVehicleModel(),
			VehicleStatus: request.GetVehicleStatus(),
			LicenseNumber: request.GetLicenseNumber(),
			Mileage:       request.GetMileage(),
		},
	)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_UpdateVehicle", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success_UpdateVehicle", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
