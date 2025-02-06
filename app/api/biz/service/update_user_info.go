package service

import (
	"context"
	"douyin_mall/api/hertz_gen/api/user"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcUser "douyin_mall/rpc/kitex_gen/user"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Run(req *user.UpdateUserInfoRequest) (resp *user.UpdateUserInfoResponse, err error) {
	ctx := h.Context
	updateUserResp, err := rpc.UserClient.UpdateUser(ctx, &rpcUser.UpdateUserReq{
		UserId:      ctx.Value(constant.UserId).(int32),
		Username:    req.Username,
		Email:       req.Email,
		Sex:         req.Sex,
		Description: req.Description,
		Avatar:      req.Avatar,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "更新用户信息失败: %v", err)
		return nil, errors.New("更新用户信息失败，请稍后再试")
	}
	return &user.UpdateUserInfoResponse{
		StatusCode: updateUserResp.StatusCode,
		StatusMsg:  updateUserResp.StatusMsg,
	}, nil
}
