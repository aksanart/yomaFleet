package usecase

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksanart/vehicle/contract"
	"github.com/aksanart/vehicle/pkg/config"
	"github.com/aksanart/vehicle/pkg/logger"
	"github.com/aksanart/vehicle/repository"
	"github.com/aksanart/vehicle/repository/repomock"
	"go.uber.org/mock/gomock"
)

func TestUseCase_UpdateVehicle(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request *contract.UpdateVehicleReq
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
				request: &contract.UpdateVehicleReq{
					Id:            "1",
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
				},
			},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().Update(gomock.Any(),gomock.Any(),gomock.Any()).Return(nil)
			},
		},
		{
			name: "err-validate",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateVehicleReq{	},
			},
			wantResponse:nil,
			wantErr: true,
			mock: func() {
			},
		},
		{
			name: "err-db",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateVehicleReq{
					Id:            "1",
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
				},
			},
			wantResponse: nil,
			wantErr: true,
			mock: func() {
				mockMongo.EXPECT().Update(gomock.Any(),gomock.Any(),gomock.Any()).Return(errors.New("error"))
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
			gotResponse, err := u.UpdateVehicle(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.UpdateVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.UpdateVehicle() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
