package usecase

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksanart/tracker_service/contract"
	"github.com/aksanart/tracker_service/model"
	"github.com/aksanart/tracker_service/pkg/config"
	"github.com/aksanart/tracker_service/pkg/logger"
	"github.com/aksanart/tracker_service/repository"
	"github.com/aksanart/tracker_service/repository/repomock"
	"go.uber.org/mock/gomock"
)

func TestUseCase_DetailTracker(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request *contract.IDTrackerReq
	}
	tests := []struct {
		name         string
		u            *UseCase
		args         args
		wantResponse *contract.DetailTrackerResponse
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.IDTrackerReq{
					Id: "123",
				},
			},
			wantResponse: &contract.DetailTrackerResponse{
				Code:    http.StatusOK,
				Message: "success",
				Data: &contract.DetailTrackerResponse_Data{
					Id:        "1",
					VehicleId: "1",
					Location: []*contract.Location{{
						Latitude:  1,
						Longitude: 1,
					}},
					CreatedAt: 0,
					UpdatedAt: 0,
				},
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(),gomock.Any()).Return(&model.Tracker{
					ID:        "1",
					VehicleID: "1",
					Location:  []model.Location{{Latitude: 1,Longitude: 1}},
					CreatedAt: 0,
					UpdatedAt: 0,
				},nil)
			},
		},
		{
			name: "err-db",
			u:    &UseCase{},
			args: args{
				ctx: ctx,
				request: &contract.IDTrackerReq{
				},
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
				request: &contract.IDTrackerReq{
					Id: "123",
				},
			},
			wantResponse:nil,
			wantErr: true,
			mock: func() {
				mockMongo.EXPECT().FindByVehicleId(gomock.Any(),gomock.Any()).Return(nil,errors.New("error"))
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
			gotResponse, err := u.DetailTracker(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.DetailTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.DetailTracker() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
