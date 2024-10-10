package usecase

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aksanart/tracker_service/contract"
	"github.com/aksanart/tracker_service/pkg/config"
	"github.com/aksanart/tracker_service/pkg/logger"
	"github.com/aksanart/tracker_service/repository"
	"github.com/aksanart/tracker_service/repository/repomock"
	"go.uber.org/mock/gomock"
)

func TestUseCase_DeleteTracker(t *testing.T) {
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
		wantResponse *contract.DefaultResponse
		wantErr      bool
		mock         func()
	}{
		{
			name:         "success",
			u:            &UseCase{},
			args:         args{
				ctx:     ctx,
				request: &contract.IDTrackerReq{
					Id: "1",
				},
			},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr:      false,
			mock: func() {
				mockMongo.EXPECT().Delete(gomock.Any(),gomock.Any()).Return(nil)
			},
		},
		{
			name:         "err-validate",
			u:            &UseCase{},
			args:         args{
				ctx:     ctx,
				request: &contract.IDTrackerReq{				},
			},
			wantResponse: nil,
			wantErr:      true,
			mock: func() {
			},
		},
		{
			name:         "err-db",
			u:            &UseCase{},
			args:         args{
				ctx:     ctx,
				request: &contract.IDTrackerReq{
					Id: "1",
				},
			},
			wantResponse:nil,
			wantErr:      true,
			mock: func() {
				mockMongo.EXPECT().Delete(gomock.Any(),gomock.Any()).Return(errors.New("error"))
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
			gotResponse, err :=u.DeleteTracker(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.DeleteTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.DeleteTracker() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
