package service

import (
	"context"
	user "douyin_mall/api/hertz_gen/api/user"
	"douyin_mall/api/infra/rpc"
	rpcuser "douyin_mall/rpc/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	client := rpc.UserClient
	res, err := client.Login(h.Context, &rpcuser.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "rpc调用登录失败, err: %v", err)
		return nil, errors.New("登录失败，请稍后再试")
	}
	resp = &user.LoginResponse{
		StatusCode:   res.StatusCode,
		StatusMsg:    res.StatusMsg,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
	return
}
