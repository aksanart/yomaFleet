package usecase

import (
	"context"
	"net/http"
	"testing"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/model"
	"github.com/aksan/weplus/apigw/pkg/config"
	"github.com/aksan/weplus/apigw/pkg/constant"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/repository"
	"github.com/aksan/weplus/apigw/repository/repomock"
	"go.uber.org/mock/gomock"
)

func TestUseCase_Login(t *testing.T) {
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
		wantResponse *contract.LoginResponse
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.RegisterReq{
					Email:    "tes@tes.com",
					Password: "1234",
				},
			},
			wantResponse: &contract.LoginResponse{
				Code:    http.StatusOK,
				Message: "123",
				Data: &contract.LoginResponse_Data{
					SessionId: "1",
				},
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(&model.User{
					ID:        "1",
					Email:     "tes@tes.com",
					Password:  "$2a$10$5NqseGd6Qv.QuyOQLaUJ1OOZ6Jv1vYahAB1aoAN7sRKaGegXoJxyi",
					CreatedAt: 0,
					UpdatedAt: 0,
				},nil)
				mockRedis.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), constant.JWT_EXPIRE).Return(nil,)
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
			_, err := u.Login(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
