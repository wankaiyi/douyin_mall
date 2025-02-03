package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/rpc/kitex_gen/auth"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/infra/rpc"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		err = model.DeleteUser(tx, req.UserId)
		if err != nil {
			klog.Errorf("数据库删除用户失败，userId：%d, err：%v", req.UserId, err)
			return errors.WithStack(err)
		}
		// 删除用户的登录信息
		_, err := rpc.AuthClient.RevokeTokenByRPC(s.ctx, &auth.RevokeTokenReq{
			UserId: req.UserId,
		})
		if err != nil {
			klog.Errorf("删除用户登录信息失败，userId：%d, err：%v", req.UserId, err)
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		klog.Errorf("删除用户事务出错，userId：%d, err：%v", req.UserId, err)
		return nil, errors.WithStack(err)
	}
	return &user.DeleteUserResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
