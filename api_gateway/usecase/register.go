package usecase

import (
	"context"
	"net/http"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/model"
	uerror "github.com/aksan/weplus/apigw/pkg/error"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/pkg/util"
)

// Register implements transport.UseCaseContract.
func (u *UseCase) Register(ctx context.Context, request *contract.RegisterReq) (response *contract.DefaultResponse, err error) {
	logData := map[string]any{
		"method":  "Register",
		"request": request,
	}
	if request.GetEmail() == "" || request.GetPassword() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	if !util.IsValidEmail(request.GetEmail()){
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	data, err := u.repo.MongoDb.User.FindByEmail(ctx, request.GetEmail())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_FindByEmail", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	if data != nil {
		logData["error"] = uerror.ErrorUserRegistered.BuildError(ctx)
		logger.GetLogger().Error("err_FindByEmail", logger.ConvertMapToFields(logData)...)
		return nil, uerror.ErrorUserRegistered.BuildError(ctx)
	}
	hashPass, err := util.HashPassword(request.GetPassword())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_HashPassword", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	err = u.repo.MongoDb.User.Create(ctx, &model.User{
		Email:    request.GetEmail(),
		Password: hashPass,
	})
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_Create", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}

	logger.GetLogger().Debug("success_Register", logger.ConvertMapToFields(logData)...)
	return &contract.DefaultResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
