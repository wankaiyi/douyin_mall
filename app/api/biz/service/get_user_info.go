package service

import (
	"context"

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
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
