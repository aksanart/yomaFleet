package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/model"
	"github.com/aksan/weplus/apigw/pkg/config"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/pkg/util"
	"github.com/aksan/weplus/apigw/repository"
	"github.com/aksan/weplus/apigw/repository/repomock"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
	vehicleContract "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
)

func TestUseCase_ListVehicle(t *testing.T) {
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
		request *contract.ListVehicleReq
	}
	tests := []struct {
		name         string
		u            *UseCase
		args         args
		wantResponse *contract.ListVehicleResponse
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.ListVehicleReq{
					Offest: 0,
				},
			},
			wantResponse: &contract.ListVehicleResponse{
				Code:    http.StatusOK,
				Message: "success",
				Data: []*contract.ListVehicleResponse_Data{{
					Id:            "1",
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
				}},
			},
			wantErr: false,
			mock: func() {
				mockRedis.EXPECT().Get(ctx, fmt.Sprintf("session:%s", token)).Return(string(sessionByte))
				mockVehicle.EXPECT().ListVehicle(gomock.Any(),gomock.Any()).Return(&vehicleContract.ListVehicleResponse{
					Code:    http.StatusOK,
					Message: "success",
					Data:    []*vehicleContract.ListVehicleResponse_Data{{
						Id:            "1",
						VehicleName:   "1",
						VehicleModel:  "1",
						VehicleStatus: "1",
						Mileage:       1,
						LicenseNumber: "1",
					}},
				}, nil)
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
			gotResponse, err := u.ListVehicle(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ListVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.ListVehicle() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
