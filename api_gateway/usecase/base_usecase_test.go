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
)

func TestUseCase_HealthCheck(t *testing.T) {
	config.LoadConfigMap()
	logger.LoadLogger()
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()
	// mockRepo := repomock.NewMockRepo(ctrl)(ctrl)
	type args struct {
		ctx     context.Context
		request *contract.EmptyRequest
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
			name: "success-health",
			u:    &UseCase{},
			args: args{},
			wantResponse: &contract.DefaultResponse{
				Code:    http.StatusOK,
				Message: "success",
			},
			wantErr: false,
			mock: func() {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo:     &repository.Repository{},
				basePath: "",
			}
			gotResponse, err := u.HealthCheck(tt.args.ctx, tt.args.request)
			tt.mock()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("UseCase.HealthCheck() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
