package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/vehicle/contract"
	uerror "github.com/aksanart/vehicle/pkg/error"
	"github.com/aksanart/vehicle/pkg/logger"
)

// ListVehicle implements transport.UseCaseContract.
func (u *UseCase) ListVehicle(ctx context.Context, request *contract.ListVehicleReq) (response *contract.ListVehicleResponse, err error) {
	logData := map[string]any{
		"method":  "ListVehicle",
		"request": request,
	}
	data, err := u.repo.MongoDb.Vehicle.FindAllVehilce(ctx, int(request.GetOffest()))
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_ListVehicle", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	var datas []*contract.ListVehicleResponse_Data
	for _, v := range data {
		d := contract.ListVehicleResponse_Data{
			Id:            v.ID,
			VehicleName:   v.VehicleName,
			VehicleModel:  v.VehicleModel,
			VehicleStatus: v.VehicleStatus,
			Mileage:       v.Mileage,
			LicenseNumber: v.LicenseNumber,
		}
		datas = append(datas, &d)
	}
	logger.GetLogger().Debug("success_ListVehicle", logger.ConvertMapToFields(logData)...)
	return &contract.ListVehicleResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    datas,
	}, nil
}
