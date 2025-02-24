package service

import (
	"context"
	user "douyin_mall/api/hertz_gen/api/user"
	"douyin_mall/api/infra/rpc"
	rpcuser "douyin_mall/rpc/kitex_gen/user"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	client := rpc.UserClient
	res, err := client.Register(h.Context, &rpcuser.RegisterReq{
		Username:        req.Username,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.Password,
		Sex:             req.Sex,
		Description:     req.Description,
		Avatar:          req.Avatar,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "rpc调用register失败, req: %v, err: %v", req, err)
		return nil, errors.New("注册失败，请稍后再试")
	}
	resp = &user.RegisterResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}
	return
}
