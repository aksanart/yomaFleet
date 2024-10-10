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
	trackerContract "github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	vehicleContract "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/pkg/util"
	"github.com/aksan/weplus/apigw/repository"
	"github.com/aksan/weplus/apigw/repository/repomock"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
)

func TestUseCase_HistoryLocationVehicle(t *testing.T) {
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
		request *contract.HistoryLocationVehicleReq
	}
	tests := []struct {
		name         string
		u            *UseCase
		args         args
		wantResponse *contract.HistoryLocationVehicleResponse
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			u:    &UseCase{},
			args: args{
				ctx:     ctx,
				request: &contract.HistoryLocationVehicleReq{
					VehicleId: "1",
				},
			},
			wantResponse: &contract.HistoryLocationVehicleResponse{
				Code:    http.StatusOK,
				Message: "success",
				Data: &contract.HistoryLocationVehicleResponse_Data{
					VehicleId:     "1",
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
					Location: []*contract.Location{{
						Latitude:  0,
						Longitude: 0,
					}},
				},
			},
			wantErr: false,
			mock: func() {
				mockRedis.EXPECT().Get(ctx, fmt.Sprintf("session:%s", token)).Return(string(sessionByte))
				mockVehicle.EXPECT().DetailVehicle(gomock.Any(), gomock.Any()).Return(&vehicleContract.DetailVehicleResponse{
					Code:    http.StatusOK,
					Message: "success",
					Data: &vehicleContract.DetailVehicleResponse_Data{
						Id:            "1",
						VehicleName:   "1",
						VehicleModel:  "1",
						VehicleStatus: "1",
						Mileage:       1,
						LicenseNumber: "1",
					},
				}, nil)
				mockTracker.EXPECT().DetailTracker(gomock.Any(), gomock.Any()).Return(&trackerContract.DetailTrackerResponse{
					Code:    http.StatusOK,
					Message: "success",
					Data: &trackerContract.DetailTrackerResponse_Data{
						Id:        "1",
						VehicleId: "1",
						Location: []*trackerContract.Location{{
							Latitude:  0,
							Longitude: 0,
						}},
						CreatedAt: 0,
						UpdatedAt: 0,
					},
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
			gotResponse, err := u.HistoryLocationVehicle(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.HistoryLocationVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.HistoryLocationVehicle() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
