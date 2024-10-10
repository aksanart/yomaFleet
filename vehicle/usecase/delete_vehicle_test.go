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
 
func TestUseCase_DeleteVehicle(t *testing.T) {
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
		u            *UseCase
		args         args
		wantResponse *contract.DefaultResponse
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				request: &contract.IDVehicleReq{
					Id: "1234",
				},
			},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "err-validate",
			args: args{
				ctx:     ctx,
				request: &contract.IDVehicleReq{},
			},
			wantResponse: nil,
			wantErr:      true,
			mock:         func() {},
		},
		{
			name: "err-db",
			args: args{
				ctx: ctx,
				request: &contract.IDVehicleReq{
					Id: "1234",
				},
			},
			wantResponse: nil,
			wantErr:      true,
			mock: func() {
				mockMongo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("error"))
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
			gotResponse, err := u.DeleteVehicle(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.DeleteVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.DeleteVehicle() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
