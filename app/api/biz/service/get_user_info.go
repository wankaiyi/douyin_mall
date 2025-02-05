package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcuser "douyin_mall/rpc/kitex_gen/user"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	user "douyin_mall/api/hertz_gen/api/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserInfoService(Context context.Context, RequestContext *app.RequestContext) *GetUserInfoService {
	return &GetUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserInfoService) Run(req *user.Empty) (resp *user.GetUserInfoResponse, err error) {
	userClient := rpc.UserClient
	ctx := h.Context
	userInfo, err := userClient.GetUser(ctx, &rpcuser.GetUserReq{
		UserId: ctx.Value(constant.UserId).(int32),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "获取用户信息失败: %v", err)
		return nil, errors.New("获取用户信息失败")
	}
	return &user.GetUserInfoResponse{
		StatusCode:  0,
		StatusMsg:   constant.GetMsg(0),
		Username:    userInfo.User.Username,
		Email:       userInfo.User.Email,
		Sex:         userInfo.User.Sex,
		Description: userInfo.User.Description,
		Avatar:      userInfo.User.Avatar,
		CreatedAt:   userInfo.User.CreatedAt,
	}, nil
}
