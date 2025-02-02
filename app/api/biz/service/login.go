package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcuser "douyin_mall/rpc/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"

	user "douyin_mall/api/hertz_gen/api/user"
	"github.com/cloudwego/hertz/pkg/app"
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
		klog.Error("登录失败, err: ", err)
		return &user.LoginResponse{
			StatusCode: 500,
			StatusMsg:  "登录失败，请稍后再试",
		}, nil
	}
	resp = &user.LoginResponse{
		StatusCode:   res.StatusCode,
		StatusMsg:    res.StatusMsg,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
	return
}
