package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/tracker_service/contract"
	uerror "github.com/aksanart/tracker_service/pkg/error"
	"github.com/aksanart/tracker_service/pkg/logger"
)

// DetailTracker implements transport.UseCaseContract.
func (u *UseCase) DetailTracker(ctx context.Context, request *contract.IDTrackerReq) (response *contract.DetailTrackerResponse, err error) {
	logData := map[string]any{
		"method":  "DetailTracker",
		"request": request,
	}
	if request.GetId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	data, err := u.repo.MongoDb.Tracker.FindByVehicleId(ctx, request.GetId())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_DetailTracker", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	datarrLocation := []*contract.Location{}
	var datas contract.DetailTrackerResponse_Data
	if data != nil {
		for _, location := range data.Location {
			loc := contract.Location{
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
			}
			datarrLocation = append(datarrLocation, &loc)
		}
		datas = contract.DetailTrackerResponse_Data{
			Id:        data.ID,
			VehicleId: data.VehicleID,
			Location:  datarrLocation,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
	}
	logger.GetLogger().Debug("success_DetailTracker", logger.ConvertMapToFields(logData)...)
	return &contract.DetailTrackerResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    &datas,
	}, nil
}
