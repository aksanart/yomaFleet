package usecase

import (
	"context"
	"time"

	// "encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/model"
	"github.com/aksan/weplus/apigw/pkg/constant"
	uError "github.com/aksan/weplus/apigw/pkg/error"
	uerror "github.com/aksan/weplus/apigw/pkg/error"
	trackerContract "github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	vehicleContract "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/goinggo/mapstructure"
	"golang.org/x/exp/rand"
)

// LiveTracking implements transport.UseCaseContract.
func (u *UseCase) LiveTrackingOne(ctx context.Context, request *contract.LiveTrackingOneReq, stream contract.ApiGateway_LiveTrackingOneServer) (err error) {
	logData := map[string]any{
		"method":  "LiveTracking",
		"request": request,
	}
	_, err = u.validateSession(ctx)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err-validateSession", logger.ConvertMapToFields(logData)...)
		return err
	}
	if request.GetVehicleId() == "" {
		return uerror.ErrorValidationRequest.BuildError(ctx)
	}
	vehicle, err := u.repo.VehicleService.DetailVehicle(ctx, &vehicleContract.IDVehicleReq{Id: request.GetVehicleId()})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_redis-historyDataStream", logger.ConvertMapToFields(logData)...)
		return uError.ErrorStream.BuildError(ctx)
	}
	if vehicle.GetData() == nil {
		logData["error"] = err
		logger.GetLogger().Error("err_redis-historyDataStream", logger.ConvertMapToFields(logData)...)
		return uError.ErrorNoData.BuildError(ctx)
	}
	// this is only mocking gps from vehicle,
	// async send gps data for every 5 second to tracker-service
	go func(ctx context.Context, vehicleID string) {
		for {
			<-time.After(5 * time.Second)
			_, err := u.repo.TrackerService.UpdateTracker(ctx, &trackerContract.UpdateTrackerReq{
				Id:        vehicleID,
				VehicleId: vehicleID,
				Location: &trackerContract.Location{
					Latitude:  rand.Float64(),
					Longitude: rand.Float64(),
				},
			})
			if err != nil {
				logger.GetLogger().Error("err_goroutine-UpdateTracker", logger.Field{Key: "err", Value: err})
				return
			}
		}
	}(ctx, request.GetVehicleId())
	sName := fmt.Sprintf(constant.REDIS_KEY_STREAM, request.GetVehicleId())
	err = u.historyDataStream(ctx, sName, vehicle.GetData(), stream)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_redis-historyDataStream", logger.ConvertMapToFields(logData)...)
		return uError.ErrorStream.BuildError(ctx)
	}
	err = u.newDataStream(ctx, sName, vehicle.GetData(), stream)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_redis-newDataStream", logger.ConvertMapToFields(logData)...)
		return uError.ErrorStream.BuildError(ctx)
	}
	return nil
}

// pull history stream
func (u *UseCase) historyDataStream(ctx context.Context, sName string, vehicle *vehicleContract.DetailVehicleResponse_Data, stream contract.ApiGateway_LiveTrackingOneServer) (err error) {
	messages, err := u.repo.Redis.RangeData(ctx, sName)
	if err != nil {
		return uError.ErrorStream.BuildError(ctx)
	}
	for _, val := range messages {
		var parsed model.DtoLiveTrackingResponse
		err = mapstructure.Decode(val.Values, &parsed)
		if err != nil {
			return uError.ErrorStream.BuildError(ctx)
		}
		if err := u.sendLiveTrackingData(&parsed, vehicle, stream); err != nil {
			return uError.ErrorStream.BuildError(ctx)
		}
	}
	return nil
}

// pull new stream
func (u *UseCase) newDataStream(ctx context.Context, sName string, vehicle *vehicleContract.DetailVehicleResponse_Data, stream contract.ApiGateway_LiveTrackingOneServer) (err error) {
	for {
		rStream, err := u.repo.Redis.GetStreamClient(ctx, sName)
		if err != nil {
			// stream closed or no data
			if err == redis.Nil || errors.Is(err, context.Canceled) {
				// need to delete
				err2 := u.repo.Redis.Del(context.Background(), sName)
				if err2 != nil {
					logger.GetLogger().Error("err_redis_GetStreamClient-Del", logger.Field{Key: "err2", Value: err2})
				}
				return nil
			}
			logger.GetLogger().Error("err_redis_GetStreamClient", logger.Field{Key: "err", Value: err})
			return uError.ErrorStream.BuildError(ctx)
		}
		for _, val := range rStream {
			for _, msg := range val.Messages {
				var parsed model.DtoLiveTrackingResponse
				err = mapstructure.Decode(msg.Values, &parsed)
				if err != nil {
					return err
				}
				if err := u.sendLiveTrackingData(&parsed, vehicle, stream); err != nil {
					return err
				}
			}
		}
	}
}

func (uc *UseCase) sendLiveTrackingData(parsed *model.DtoLiveTrackingResponse, vehicle *vehicleContract.DetailVehicleResponse_Data, stream contract.ApiGateway_LiveTrackingOneServer,
) error {
	logData := map[string]any{
		"method": "LiveTracking-parseAndSendLiveTrackingData",
		"parsed": parsed,
	}
	response := &contract.LiveTrackingOneResp{
		VehicleId:     vehicle.GetId(),
		VehicleName:   vehicle.GetVehicleName(),
		VehicleModel:  vehicle.GetVehicleModel(),
		VehicleStatus: vehicle.GetVehicleStatus(),
		Mileage:       vehicle.GetMileage(),
		LicenseNumber: vehicle.GetLicenseNumber(),
		Latitude:      stringToFloat(parsed.Latitude),
		Longitude:     stringToFloat(parsed.Longitude),
	}
	if err := stream.Send(response); err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_parseAndSendLiveTrackingData_stream", logger.ConvertMapToFields(logData)...)
		return err
	}
	return nil
}

func stringToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return f
}
