package usecase

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksanart/vehicle/contract"
	"github.com/aksanart/vehicle/model"
	"github.com/aksanart/vehicle/pkg/config"
	"github.com/aksanart/vehicle/pkg/logger"
	"github.com/aksanart/vehicle/repository"
	"github.com/aksanart/vehicle/repository/repomock"
	"go.uber.org/mock/gomock"
)

func TestUseCase_ListVehicle(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	ctx := context.Background()
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
			args: args{ctx: ctx, request: &contract.ListVehicleReq{
				Offest: 0,
			}},
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
				mockMongo.EXPECT().FindAllVehilce(ctx, 0).Return([]*model.Vehicle{{
					ID:            "1",
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1"}}, nil,
				)
			},
		},
		{
			name: "err-db",
			u:    &UseCase{},
			args: args{ctx: ctx, request: &contract.ListVehicleReq{
				Offest: 0,
			}},
			wantResponse: nil,
			wantErr:      true,
			mock: func() {
				mockMongo.EXPECT().FindAllVehilce(ctx, 0).Return(nil, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: &repository.Repository{
					MongoDb: repository.MongoCollections{
						Vehicle: mockMongo,
					},
					Redis: mockRedis,
				},
				basePath: "",
			}
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
