package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type GetUserRoleByIdService struct {
	ctx context.Context
} // NewGetUserRoleByIdService new GetUserRoleByIdService
func NewGetUserRoleByIdService(ctx context.Context) *GetUserRoleByIdService {
	return &GetUserRoleByIdService{ctx: ctx}
}

// Run create note info
func (s *GetUserRoleByIdService) Run(req *user.GetUserRoleByIdReq) (resp *user.GetUserRoleByIdResp, err error) {
	role, err := model.GetUserRoleById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "查询用户角色失败，userId: %d, err: %v", req.UserId, err)
		return nil, errors.WithStack(err)
	}
	return &user.GetUserRoleByIdResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Role:       role,
	}, nil
}
