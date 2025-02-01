package service

import (
	"context"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	"fmt"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	fmt.Print(s.ctx.Value("user_id"))
	userInfo, err := model.GetUserById(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &user.GetUserResp{
		StatusCode: 0,
		StatusMsg:  "success",
		User: &user.User{
			Email: userInfo.Email,
		},
	}
	return
}
