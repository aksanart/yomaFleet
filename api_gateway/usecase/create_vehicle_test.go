package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/model"
	"github.com/aksan/weplus/apigw/pkg/config"
	trackerContract "github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	vehicleContract "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/pkg/util"
	"github.com/aksan/weplus/apigw/repository"
	"github.com/aksan/weplus/apigw/repository/repomock"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
)

func TestUseCase_CreateVehicle(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	mockVehicle := repomock.NewMockVehicleServiceIface(ctrl)
	mockTracker := repomock.NewMockTrackerServiceIface(ctrl)
	token, _ := util.GenerateJwt("test@example.com")
	md := metadata.New(make(map[string]string))
	md.Append("Session-Id", token)
	sessionData := model.SessionData{Jwt: token}
	sessionByte, _ := json.Marshal(sessionData)
	ctx := metadata.NewIncomingContext(context.Background(), md)
	type args struct {
		ctx     context.Context
		request *contract.CreateVehicleReq
	}
	tests := []struct {
		name         string
		u            *UseCase
		args         args
		wantResponse *contract.DefaultResponse
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.CreateVehicleReq{
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
					Latitude:      1,
					Longitude:     1,
				},
			},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr: false,
			mock: func() {
				mockRedis.EXPECT().Get(ctx, fmt.Sprintf("session:%s", token)).Return(string(sessionByte))
				mockVehicle.EXPECT().CreateVehicle(gomock.Any(), gomock.Any()).Return(&vehicleContract.CreateVehicleResponse{
					Code:    http.StatusOK,
					Message: "success",
					Data: &vehicleContract.CreateVehicleResponse_Data{
						Id: "1",
					},
				}, nil)
				mockTracker.EXPECT().CreateTracker(gomock.Any(), gomock.Any()).Return(&trackerContract.CreateTrackerResponse{
					Id: "1",
				}, nil)
			},
		},
		{
			name: "err-tracker",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.CreateVehicleReq{
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
					Latitude:      1,
					Longitude:     1,
				},
			},
			wantResponse: nil,
			wantErr: true,
			mock: func() {
				mockRedis.EXPECT().Get(ctx, fmt.Sprintf("session:%s", token)).Return(string(sessionByte))
				mockVehicle.EXPECT().CreateVehicle(gomock.Any(), gomock.Any()).Return(&vehicleContract.CreateVehicleResponse{
					Code:    http.StatusOK,
					Message: "success",
					Data: &vehicleContract.CreateVehicleResponse_Data{
						Id: "1",
					},
				}, nil)
				mockTracker.EXPECT().CreateTracker(gomock.Any(), gomock.Any()).Return(&trackerContract.CreateTrackerResponse{
					Id: "1",
				}, errors.New("error"))
			},
		},
		{
			name: "err-vehicle",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.CreateVehicleReq{
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
					Latitude:      1,
					Longitude:     1,
				},
			},
			wantResponse: nil,
			wantErr: true,
			mock: func() {
				mockRedis.EXPECT().Get(ctx, fmt.Sprintf("session:%s", token)).Return(string(sessionByte))
				mockVehicle.EXPECT().CreateVehicle(gomock.Any(), gomock.Any()).Return(&vehicleContract.CreateVehicleResponse{
					Code:    http.StatusOK,
					Message: "success",
					Data: &vehicleContract.CreateVehicleResponse_Data{
						Id: "1",
					},
				}, errors.New("error"))
			},
		},
		{
			name: "err-vehicle",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.CreateVehicleReq{
					// VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
					Latitude:      1,
					Longitude:     1,
				},
			},
			wantResponse: nil,
			wantErr: true,
			mock: func() {
				mockRedis.EXPECT().Get(ctx, fmt.Sprintf("session:%s", token)).Return(string(sessionByte))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		u := &UseCase{
			repo: &repository.Repository{
				MongoDb:        repository.MongoCollections{User: mockMongo},
				Redis:          mockRedis,
				VehicleService: mockVehicle,
				TrackerService: mockTracker,
			},
			basePath: "",
		}
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := u.CreateVehicle(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CreateVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.CreateVehicle() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
