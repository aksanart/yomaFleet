package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/vehicle/contract"
	uerror "github.com/aksanart/vehicle/pkg/error"
	"github.com/aksanart/vehicle/pkg/logger"
)

// DetailVehicle implements transport.UseCaseContract.
func (u *UseCase) DetailVehicle(ctx context.Context, request *contract.IDVehicleReq) (response *contract.DetailVehicleResponse, err error) {
	logData := map[string]any{
		"method":  "DetailVehicle",
		"request": request,
	}
	if request.GetId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	data, err := u.repo.MongoDb.Vehicle.FindById(ctx, request.GetId())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_DetailVehicle", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success_DetailVehicle", logger.ConvertMapToFields(logData)...)
	if data == nil {
		return &contract.DetailVehicleResponse{
			Code:    http.StatusOK,
			Message: "success", Data: nil}, nil
	}
	return &contract.DetailVehicleResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data: &contract.DetailVehicleResponse_Data{
			Id:            data.ID,
			VehicleName:   data.VehicleName,
			VehicleModel:  data.VehicleModel,
			VehicleStatus: data.VehicleStatus,
			Mileage:       data.Mileage,
			LicenseNumber: data.LicenseNumber,
		},
	}, nil
}
