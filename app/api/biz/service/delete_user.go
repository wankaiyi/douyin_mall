package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcUser "douyin_mall/rpc/kitex_gen/user"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	user "douyin_mall/api/hertz_gen/api/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUserService(Context context.Context, RequestContext *app.RequestContext) *DeleteUserService {
	return &DeleteUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUserService) Run(req *user.Empty) (resp *user.DeleteUserResponse, err error) {
	ctx := h.Context
	deleteUserResp, err := rpc.UserClient.DeleteUser(ctx, &rpcUser.DeleteUserReq{
		UserId: ctx.Value(constant.UserId).(int32),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "删除用户失败: %v", err)
		return nil, errors.New("删除用户失败，请稍后再试")
	}
	return &user.DeleteUserResponse{
		StatusCode: deleteUserResp.StatusCode,
		StatusMsg:  deleteUserResp.StatusMsg,
	}, nil
}
