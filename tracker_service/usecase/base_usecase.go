package usecase

import (
	"context"
	"net/http"

	"github.com/aksanart/tracker_service/contract"
	"github.com/aksanart/tracker_service/repository"
)

// var useCasePointer transport.UseCaseContract
var useCasePointer *UseCase

type UseCase struct {
	repo     *repository.Repository
	basePath string
}

// HealthCheck implements transport.UseCaseContract.
func (u *UseCase) HealthCheck(ctx context.Context, request *contract.EmptyRequest) (response *contract.DefaultResponse, err error) {
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}

func NewUsecase(repoIn *repository.Repository, basePath string) *UseCase {
	if useCasePointer == nil {
		useCasePointer = &UseCase{repo: repoIn, basePath: basePath}
	}
	return useCasePointer
}
