package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/tracker_service/contract"
	uerror "github.com/aksanart/tracker_service/pkg/error"
	"github.com/aksanart/tracker_service/pkg/logger"
)

// ListTracker implements transport.UseCaseContract.
func (u *UseCase) ListTracker(ctx context.Context, request *contract.EmptyRequest) (response *contract.ListTrackerResponse, err error) {
	logData := map[string]any{
		"method":  "ListTracker",
		"request": request,
	}
	data, err := u.repo.MongoDb.Tracker.FindAllTracker(ctx)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_ListTracker", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	var datas []*contract.ListTrackerResponse_Data
	for _, v := range data {
		arrLocation := []*contract.Location{}
		for _, location := range v.Location {
			loc := contract.Location{
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
			}
			arrLocation = append(arrLocation, &loc)
		}
		d := contract.ListTrackerResponse_Data{
			Id:        v.ID,
			VehicleId: v.VehicleID,
			Location:  arrLocation,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		datas = append(datas, &d)
	}
	logger.GetLogger().Debug("success_ListTracker", logger.ConvertMapToFields(logData)...)
	return &contract.ListTrackerResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    datas,
	}, nil
}
