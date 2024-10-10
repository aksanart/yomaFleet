package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/model"
	"github.com/aksan/weplus/apigw/pkg/constant"
	uerror "github.com/aksan/weplus/apigw/pkg/error"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/pkg/util"
	"github.com/google/uuid"
)

// Login implements transport.UseCaseContract.
func (u *UseCase) Login(ctx context.Context, request *contract.RegisterReq) (response *contract.LoginResponse, err error) {
	logData := map[string]any{
		"method":  "Login",
		"request": request,
	}
	if request.GetEmail() == "" || request.GetPassword() == "" {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	if !util.IsValidEmail(request.GetEmail()) {
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	data, err := u.repo.MongoDb.User.FindByEmail(ctx, request.GetEmail())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_FindByEmail", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	if data == nil {
		logData["error"] = errors.New("no user found")
		logger.GetLogger().Error("err_FindByEmail", logger.ConvertMapToFields(logData)...)
		return nil, uerror.ErrorNotFound.BuildError(ctx)
	}
	hashPass, err := util.HashPassword(request.GetPassword())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_HashPassword", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	cekPass := util.CheckPasswordHash(request.GetPassword(), hashPass)
	if !cekPass {
		logData["error"] = uerror.ErrorValidationRequest.BuildError(ctx)
		logger.GetLogger().Error("err_FindByEmail", logger.ConvertMapToFields(logData)...)
		return nil, uerror.ErrorValidationRequest.BuildError(ctx)
	}
	tokenJwt, err := util.GenerateJwt(request.GetEmail())
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_GenerateJwt", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	//set session id to redis
	sessionID := uuid.NewString()
	sessionData := model.SessionData{
		Jwt: tokenJwt,
	}
	sessionKey := fmt.Sprintf("session:%s", sessionID)
	sessionByte, err := json.Marshal(sessionData)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_Marshal", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	err = u.repo.Redis.Set(ctx, sessionKey, sessionByte, constant.JWT_EXPIRE)
	if err != nil {
		logData["error"] = err
		logger.GetLogger().Error("err_Redis.Set", logger.ConvertMapToFields(logData)...)
		return nil, uerror.InternalServerError.BuildError(ctx)
	}
	logger.GetLogger().Debug("success", logger.ConvertMapToFields(logData)...)
	return &contract.LoginResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    &contract.LoginResponse_Data{SessionId: sessionID},
	}, nil
}
