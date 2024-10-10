package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aksanart/tracker_service/contract"
	"github.com/aksanart/tracker_service/model"
	"github.com/aksanart/tracker_service/pkg/constant"
	uerror "github.com/aksanart/tracker_service/pkg/error"
	"github.com/aksanart/tracker_service/pkg/logger"
)

// UpdateTracker implements transport.UseCaseContract.
func (u *UseCase) UpdateTracker(ctx context.Context, request *contract.UpdateTrackerReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "UpdateTracker",
		"request": request,
	}
	if request.GetVehicleId() == "" || request.GetLocation() == nil {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	data, err := u.repo.MongoDb.Tracker.FindByVehicleId(ctx, request.GetVehicleId())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_FindByVehicleId", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	if data != nil {
		err = u.repo.MongoDb.Tracker.UpdateLocation(ctx, request.GetVehicleId(), &model.Tracker{
			Location: []model.Location{{Latitude: request.GetLocation().GetLatitude(), Longitude: request.GetLocation().GetLongitude()}},
		})
		if err != nil {
			logData["error"] = err
			logger.GetLogger().Error("err_UpdateTracker", logger.ConvertMapToFields(logData)...)
			return nil, uerror.InternalServerError.BuildError(ctx)
		}
	} else {
		err = u.repo.MongoDb.Tracker.Create(ctx,
			&model.Tracker{
				VehicleID: request.GetVehicleId(),
				Location:  []model.Location{{Latitude: request.GetLocation().GetLatitude(), Longitude: request.GetLocation().GetLongitude()}},
			})
		if err != nil {
			logData["error"] = err
			logger.GetLogger().Error("err_UpdateTracker", logger.ConvertMapToFields(logData)...)
			return nil, uerror.InternalServerError.BuildError(ctx)
		}
	}
	dataStream := map[string]interface{}{
		"vehicle_id": request.GetVehicleId(),
		"latitude":   request.GetLocation().GetLatitude(),
		"longitude":  request.GetLocation().GetLongitude(),
	}
	keyStream := fmt.Sprintf(constant.REDIS_KEY_STREAM, request.GetVehicleId())
	err = u.repo.Redis.SetStreamData(ctx, keyStream, dataStream)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_SetStreamData", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	// delete tracking data, if not used
	err = u.repo.Redis.ExpireStream(ctx, keyStream, constant.REDIS_EXPIRE_STREAM)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_ExpireStream", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success_UpdateTracker", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
