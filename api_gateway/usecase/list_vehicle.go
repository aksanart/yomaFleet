package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/model"
	"github.com/aksan/weplus/apigw/pkg/constant"
	uerror "github.com/aksan/weplus/apigw/pkg/error"
	vehicleContract "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/pkg/util"
	"google.golang.org/grpc/metadata"
)

// ListVehicle implements transport.UseCaseContract.
func (u *UseCase) ListVehicle(ctx context.Context, request *contract.ListVehicleReq) (response *contract.ListVehicleResponse, err error) {
	logData := map[string]any{
		"method":  "ListVehicle",
		"request": request,
	}
	_, err = u.validateSession(ctx)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-validateSession", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	vehicleData, err := u.repo.VehicleService.ListVehicle(ctx, &vehicleContract.ListVehicleReq{
		Offest: request.GetOffest(),
	})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-validateSession", logger.ConvertMapToFields(logData)...)
		return nil, err
	}
	var datas []*contract.ListVehicleResponse_Data
	for _, v := range vehicleData.GetData() {
		data := contract.ListVehicleResponse_Data{
			Id:            v.GetId(),
			VehicleName:   v.GetVehicleName(),
			VehicleModel:  v.GetVehicleModel(),
			VehicleStatus: v.GetVehicleStatus(),
			Mileage:       v.GetMileage(),
			LicenseNumber: v.GetLicenseNumber(),
		}
		datas = append(datas, &data)
	}
	logData["vehicleData"] = vehicleData
	logger.GetLogger().Debug("success", logger.ConvertMapToFields(logData)...)
	return &contract.ListVehicleResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    datas,
	}, nil
}

func (u *UseCase) validateSession(ctx context.Context) (email string, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	sessionHeader := md.Get(constant.SESSION_HEADER)
	if len(sessionHeader) == 0 {
		return "", uerror.ErrorRequiredSession.BuildError(ctx)
	}
	var sessionData model.SessionData
	sessionKey := fmt.Sprintf("session:%s", sessionHeader[0])
	dataRedis := u.repo.Redis.Get(ctx, sessionKey)
	if dataRedis == "" {
		return "", uerror.ErrorSessionExpired.BuildError(ctx)
	}
	err = json.Unmarshal([]byte(dataRedis), &sessionData)
	if err != nil {
		return "", err
	}
	data, err := util.ValidateToken(sessionData.Jwt)
	if err != nil {
		return data.Email, err
	}
	return data.Email, err
}
