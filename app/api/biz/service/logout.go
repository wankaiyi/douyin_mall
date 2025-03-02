package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	"douyin_mall/rpc/kitex_gen/auth"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	user "douyin_mall/api/hertz_gen/api/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *user.Empty) (resp *user.LogoutResponse, err error) {
	ctx := h.Context
	logOutResp, err := rpc.AuthClient.RevokeTokenByRPC(h.Context, &auth.RevokeTokenReq{
		UserId: ctx.Value(constant.UserId).(int32),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用登出接口失败: %v", err)
		return nil, errors.New("登出失败，请稍后再试")
	}
	return &user.LogoutResponse{
		StatusCode: logOutResp.StatusCode,
		StatusMsg:  logOutResp.StatusMsg,
	}, nil
}
