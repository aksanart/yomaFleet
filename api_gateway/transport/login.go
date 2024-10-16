// Code generated by proto-gen-svc-transport. DO NOT EDIT.
// Folder	: "github.com/aksan/weplus/apigw/contract"
// File		: "contract/grpc.proto"

package transport

import (
	"context"
	"github.com/aksan/weplus/apigw/contract"
)

func (t transport) Login(ctx context.Context, request *contract.RegisterReq) (response *contract.LoginResponse, err error) {
	// * Bridging to usecase related method
	return t.UseCase.Login(ctx, request)
}
