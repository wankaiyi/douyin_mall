package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/rpc/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"

	user "douyin_mall/api/hertz_gen/api/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type RefreshTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRefreshTokenService(Context context.Context, RequestContext *app.RequestContext) *RefreshTokenService {
	return &RefreshTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *RefreshTokenService) Run(req *user.Empty) (resp *user.LoginResponse, err error) {
	authClient := rpc.AuthClient
	refreshTokenResp, err := authClient.RefreshTokenByRPC(h.Context, &auth.RefreshTokenReq{
		RefreshToken: string(h.RequestContext.GetHeader("refresh_token")),
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "刷新token失败，error: %v", err)
		return nil, errors.New("刷新token失败")
	}
	return &user.LoginResponse{
		AccessToken:  refreshTokenResp.AccessToken,
		RefreshToken: refreshTokenResp.RefreshToken,
		StatusCode:   refreshTokenResp.StatusCode,
		StatusMsg:    refreshTokenResp.StatusMsg,
	}, nil
}
