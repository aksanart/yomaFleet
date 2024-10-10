package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/tracker_service/contract"
	uerror "github.com/aksanart/tracker_service/pkg/error"
	"github.com/aksanart/tracker_service/pkg/logger"
)

// DeleteTracker implements transport.UseCaseContract.
func (u *UseCase) DeleteTracker(ctx context.Context, request *contract.IDTrackerReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "DeleteTracker",
		"request": request,
	}
	if request.GetId() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	err = u.repo.MongoDb.Tracker.Delete(ctx, request.GetId())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_DeleteTracker", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success_DeleteTracker", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
