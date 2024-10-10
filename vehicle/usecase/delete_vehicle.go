package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/vehicle/contract"
	uerror "github.com/aksanart/vehicle/pkg/error"
	"github.com/aksanart/vehicle/pkg/logger"
)

// DeleteVehicle implements transport.UseCaseContract.
func (u *UseCase) DeleteVehicle(ctx context.Context, request *contract.IDVehicleReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "DeleteVehicle",
		"request": request,
	}
	if request.GetId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	err = u.repo.MongoDb.Vehicle.Delete(ctx, request.GetId())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_DeleteVehicle", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success_DeleteVehicle", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
