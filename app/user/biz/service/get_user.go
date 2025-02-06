package service

import (
	"context"
	"douyin_mall/common/utils"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	ctx := s.ctx
	userInfo, err := model.GetUserById(mysql.DB, ctx, req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, "查询用户信息失败, error: %v", err)
		return nil, errors.WithStack(err)
	}
	resp = &user.GetUserResp{
		StatusCode: 0,
		StatusMsg:  "success",
		User: &user.User{
			Username:    userInfo.Username,
			Email:       userInfo.Email,
			Sex:         model.SexToString(userInfo.Sex),
			Description: userInfo.Description,
			Avatar:      userInfo.Avatar,
			CreatedAt:   utils.GetFormattedDateTime(userInfo.CreatedAt),
		},
	}
	return
}
