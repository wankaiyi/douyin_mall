package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcuser "douyin_mall/rpc/kitex_gen/user"

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
		UserId: ctx.Value("user_id").(int32),
	})
	if err != nil {
		return nil, err
	}
	return &user.GetUserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Email:      userInfo.User.Email,
	}, nil
}
