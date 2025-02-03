package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	if err = model.UpdateUser(mysql.DB, s.ctx, req.UserId, req.Username, req.Email, req.Sex, req.Description, req.Avatar); err != nil {
		klog.Errorf("更新用户信息失败: req: %v, err: %v", req, err)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			resp.StatusCode = 1008
			resp.StatusMsg = constant.GetMsg(1008)
			return resp, nil
		}
		return nil, errors.WithStack(err)
	}
	return &user.UpdateUserResp{StatusCode: 0, StatusMsg: constant.GetMsg(0)}, nil
}
