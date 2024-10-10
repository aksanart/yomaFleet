package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/vehicle/contract"
	"github.com/aksanart/vehicle/model"
	uerror "github.com/aksanart/vehicle/pkg/error"
	"github.com/aksanart/vehicle/pkg/logger"
)

// CreateVehicle implements transport.UseCaseContract.
func (u *UseCase) CreateVehicle(ctx context.Context, request *contract.CreateVehicleReq) (response *contract.CreateVehicleResponse, err error) {
	logData := map[string]any{
		"method":  "CreateVehicle",
		"request": request,
	}
	if request.GetVehicleName() == "" || request.GetVehicleModel() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	id, err := u.repo.MongoDb.Vehicle.Create(
		ctx,
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
		logger.GetLogger().Error("err_CreateVehicle", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success_CreateVehicle", logger.ConvertMapToFields(logData)...)
	return &contract.CreateVehicleResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data: &contract.CreateVehicleResponse_Data{
			Id: id,
		},
	}, nil
}
