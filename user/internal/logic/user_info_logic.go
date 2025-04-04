package logic

import (
	"context"
	"fmt"

	"grpc_demo_server/user/internal/svc"
	"grpc_demo_server/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	// 在这里编写您的业务逻辑
	return &user.UserInfoResp{
		Name: fmt.Sprintf("Hello, %s!", in.Name),
	}, nil
}
