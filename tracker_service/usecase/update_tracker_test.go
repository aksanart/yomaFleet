package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksanart/tracker_service/contract"
	"github.com/aksanart/tracker_service/model"
	"github.com/aksanart/tracker_service/pkg/config"
	"github.com/aksanart/tracker_service/pkg/constant"
	"github.com/aksanart/tracker_service/pkg/logger"
	"github.com/aksanart/tracker_service/repository"
	"github.com/aksanart/tracker_service/repository/repomock"
	"go.uber.org/mock/gomock"
)

func TestUseCase_UpdateTracker(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request *contract.UpdateTrackerReq
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
			name: "success-insert",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
					Id:        "1",
					VehicleId: "1",
					Location: &contract.Location{
						Latitude:  1,
						Longitude: 1,
					},
				},
			},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(), gomock.Any()).Return(nil, nil)
				mockMongo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().SetStreamData(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().ExpireStream(gomock.Any(), fmt.Sprintf(constant.REDIS_KEY_STREAM, "1"), constant.REDIS_EXPIRE_STREAM).Return(nil)
			},
		},
		{
			name: "success-update",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
					Id:        "1",
					VehicleId: "1",
					Location: &contract.Location{
						Latitude:  1,
						Longitude: 1,
					},
				},
			},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(), gomock.Any()).Return(&model.Tracker{
					ID:        "1",
					VehicleID: "1",
					Location:  []model.Location{},
					CreatedAt: 0,
					UpdatedAt: 0,
				}, nil)
				mockMongo.EXPECT().UpdateLocation(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().SetStreamData(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().ExpireStream(gomock.Any(), fmt.Sprintf(constant.REDIS_KEY_STREAM, "1"), constant.REDIS_EXPIRE_STREAM).Return(nil)
			},
		},
		{
			name: "err-validate",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
				},
			},
			wantResponse:nil,
			wantErr: true,
			mock: func() {
			},
		},
		{
			name: "err-find",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
					Id:        "1",
					VehicleId: "1",
					Location: &contract.Location{
						Latitude:  1,
						Longitude: 1,
					},
				},
			},
			wantResponse:nil,
			wantErr: true,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
		{
			name: "err-insert",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
					Id:        "1",
					VehicleId: "1",
					Location: &contract.Location{
						Latitude:  1,
						Longitude: 1,
					},
				},
			},
			wantResponse:nil,
			wantErr: true,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(), gomock.Any()).Return(nil, nil)
				mockMongo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
		},
		{
			name: "err-update",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
					Id:        "1",
					VehicleId: "1",
					Location: &contract.Location{
						Latitude:  1,
						Longitude: 1,
					},
				},
			},
			wantResponse:nil,
			wantErr: true,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(), gomock.Any()).Return(&model.Tracker{
					ID:        "1",
					VehicleID: "1",
					Location:  []model.Location{},
					CreatedAt: 0,
					UpdatedAt: 0,
				}, nil)
				mockMongo.EXPECT().UpdateLocation(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
		},
		{
			name: "err-setstream",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
					Id:        "1",
					VehicleId: "1",
					Location: &contract.Location{
						Latitude:  1,
						Longitude: 1,
					},
				},
			},
			wantResponse: nil,
			wantErr: true,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(), gomock.Any()).Return(&model.Tracker{
					ID:        "1",
					VehicleID: "1",
					Location:  []model.Location{},
					CreatedAt: 0,
					UpdatedAt: 0,
				}, nil)
				mockMongo.EXPECT().UpdateLocation(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().SetStreamData(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
		},
		{
			name: "err-ExpireStream",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.UpdateTrackerReq{
					Id:        "1",
					VehicleId: "1",
					Location: &contract.Location{
						Latitude:  1,
						Longitude: 1,
					},
				},
			},
			wantResponse: nil,
			wantErr: true,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(), gomock.Any()).Return(&model.Tracker{
					ID:        "1",
					VehicleID: "1",
					Location:  []model.Location{},
					CreatedAt: 0,
					UpdatedAt: 0,
				}, nil)
				mockMongo.EXPECT().UpdateLocation(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().SetStreamData(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().ExpireStream(gomock.Any(), fmt.Sprintf(constant.REDIS_KEY_STREAM, "1"), constant.REDIS_EXPIRE_STREAM).Return(errors.New("errors"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		u := &UseCase{
			repo: &repository.Repository{
				MongoDb: repository.MongoCollections{
					Tracker: mockMongo,
				},
				Redis: mockRedis,
			},
			basePath: "",
		}
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := u.UpdateTracker(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.UpdateTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.UpdateTracker() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
