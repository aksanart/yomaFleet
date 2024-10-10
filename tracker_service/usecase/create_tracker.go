package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/tracker_service/contract"
	"github.com/aksanart/tracker_service/model"
	uerror "github.com/aksanart/tracker_service/pkg/error"
	"github.com/aksanart/tracker_service/pkg/logger"
	"github.com/google/uuid"
)

// CreateTracker implements transport.UseCaseContract.
func (u *UseCase) CreateTracker(ctx context.Context, request *contract.CreateTrackerReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "CreateTracker",
		"request": request,
	}
	if request.GetVehicleId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	var locations []model.Location
	for _, v := range request.GetLocation() {
		loc := model.Location{
			Latitude:  v.GetLatitude(),
			Longitude: v.GetLongitude(),
		}
		locations = append(locations, loc)
	}
	err = u.repo.MongoDb.Tracker.Create(
		ctx,
		&model.Tracker{
			ID:        uuid.NewString(),
			VehicleID: request.GetVehicleId(),
			Location:  locations,
		},
	)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_CreateTracker", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success_CreateTracker", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
