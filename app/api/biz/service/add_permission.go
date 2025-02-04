package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/rpc/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"

	user "douyin_mall/api/hertz_gen/api/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddPermissionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddPermissionService(Context context.Context, RequestContext *app.RequestContext) *AddPermissionService {
	return &AddPermissionService{RequestContext: RequestContext, Context: Context}
}

func (h *AddPermissionService) Run(req *user.AddPermissionRequest) (resp *user.Empty, err error) {
	_, err = rpc.AuthClient.AddPermission(h.Context, &auth.AddPermissionReq{
		Role:   req.Role,
		Path:   req.Path,
		Method: req.Method,
	})
	if err != nil {
		klog.Errorf("添加权限失败，req: %v, err: %v", req, err)
		return nil, errors.WithStack(err)
	}
	return &user.Empty{}, nil
}
