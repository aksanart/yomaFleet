package usecase

import (
	"context"
	"net/http"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/repository"
)

// var useCasePointer transport.UseCaseContract
var useCasePointer *UseCase

type UseCase struct {
	repo     *repository.Repository
	basePath string
}

func NewUsecase(repoIn *repository.Repository, basePath string) *UseCase {
	if useCasePointer == nil {
		useCasePointer = &UseCase{repo: repoIn, basePath: basePath}
	}
	return useCasePointer
}

// HealthCheck implements transport.UseCaseContract.
func (u *UseCase) HealthCheck(ctx context.Context, request *contract.EmptyRequest) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "Login",
		"request": request,
	}
	logger.GetLogger().Debug("success", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
