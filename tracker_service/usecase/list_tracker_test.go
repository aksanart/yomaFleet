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

func TestUseCase_ListTracker(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMongo := repomock.NewMockMongoInterface(ctrl)
	mockRedis := repomock.NewMockRedisInterface(ctrl)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request *contract.EmptyRequest
	}
	tests := []struct {
		name         string
		u            *UseCase
		args         args
		wantResponse *contract.ListTrackerResponse
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			u:    &UseCase{},
			args: args{
				ctx:     ctx,
				request: &contract.EmptyRequest{},
			},
			wantResponse: &contract.ListTrackerResponse{
				Code:    http.StatusOK,
				Message: "success",
				Data: []*contract.ListTrackerResponse_Data{{
					Id:        "1",
					VehicleId: "1",
					Location: []*contract.Location{{
						Latitude:  1,
						Longitude: 1,
					}},
					CreatedAt: 0,
					UpdatedAt: 0,
				}},
			},
			wantErr: false,
			mock: func() {
				mockMongo.EXPECT().FindAllTracker(ctx).Return([]*model.Tracker{{
					ID:        "1",
					VehicleID: "1",
					Location: []model.Location{{
						Latitude:  1,
						Longitude: 1,
					}},
					CreatedAt: 0,
					UpdatedAt: 0,
				}}, nil)
			},
		},
		{
			name: "err-db",
			u:    &UseCase{},
			args: args{
				ctx:     ctx,
				request: &contract.EmptyRequest{},
			},
			wantResponse: nil,
			wantErr:      true,
			mock: func() {
				mockMongo.EXPECT().FindAllTracker(ctx).Return(nil, errors.New("error"))
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
			gotResponse, err := u.ListTracker(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ListTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.ListTracker() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
