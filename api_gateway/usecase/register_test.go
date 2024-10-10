package usecase

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/pkg/config"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/repository"
	"github.com/aksan/weplus/apigw/repository/repomock"
	"go.uber.org/mock/gomock"
)

func TestUseCase_Register(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	mockVehicle := repomock.NewMockVehicleServiceIface(ctrl)
	mockTracker := repomock.NewMockTrackerServiceIface(ctrl)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request *contract.RegisterReq
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
				request: &contract.RegisterReq{
					Email:    "test@example.com",
					Password: "1234",
				},
			},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(nil,nil)
				mockMongo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
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
			gotResponse, err := u.Register(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.Register() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
