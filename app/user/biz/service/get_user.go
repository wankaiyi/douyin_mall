package service

import (
	"context"
	user "douyin_mall/user/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	// Finish your business logic.

	return
}
