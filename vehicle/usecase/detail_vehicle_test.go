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

func TestUseCase_DetailVehicle(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request *contract.IDVehicleReq
	}
	tests := []struct {
		name         string
		args         args
		wantResponse *contract.DetailVehicleResponse
		wantErr      bool
		mock         func()
	}{
		{
			name:         "success",
			args:         args{
				ctx:     ctx,
				request: &contract.IDVehicleReq{
					Id: "123",
				},
			},
			wantResponse: &contract.DetailVehicleResponse{
				Code:    http.StatusOK,
				Message: "success",
				Data:    &contract.DetailVehicleResponse_Data{
					Id:            "1",
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					Mileage:       1,
					LicenseNumber: "1",
				},
			},
			wantErr:      false,
			mock: func() {
				mockMongo.EXPECT().FindById(gomock.Any(),gomock.Any()).Return(&model.Vehicle{
					ID:            "1",
					VehicleName:   "1",
					VehicleModel:  "1",
					VehicleStatus: "1",
					LicenseNumber: "1",
					Mileage:       1,
					CreatedAt:     123,
					UpdatedAt:     123,
				},nil)
			},
		},
		{
			name:         "err-validate",
			args:         args{
				ctx:     ctx,
				request: &contract.IDVehicleReq{
				},
			},
			wantResponse: nil,
			wantErr:      true,
			mock: func() {
			},
		},
		{
			name:         "err-db",
			args:         args{
				ctx:     ctx,
				request: &contract.IDVehicleReq{
					Id: "123",
				},
			},
			wantResponse:nil,
			wantErr:      true,
			mock: func() {
				mockMongo.EXPECT().FindById(gomock.Any(),gomock.Any()).Return(nil,errors.New("error"))
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
			gotResponse, err := u.DetailVehicle(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.DetailVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.DetailVehicle() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
