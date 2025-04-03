// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"grpc_demo_server/user/internal/logic"
	"grpc_demo_server/user/internal/svc"
	"grpc_demo_server/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoReq) (*user.UserInfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}
