package service

import (
	"context"
	"douyin_mall/auth/conf"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/common/constant"
)

type CheckIfUserBannedService struct {
	ctx context.Context
} // NewCheckIfUserBannedService new CheckIfUserBannedService
func NewCheckIfUserBannedService(ctx context.Context) *CheckIfUserBannedService {
	return &CheckIfUserBannedService{ctx: ctx}
}

// Run create note info
func (s *CheckIfUserBannedService) Run(req *auth.CheckIfUserBannedReq) (resp *auth.CheckIfUserBannedResp, err error) {
	_, exist := conf.BannedUserList[req.UserId]
	resp = &auth.CheckIfUserBannedResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		IsBanned:   exist,
	}
	return resp, nil
}
