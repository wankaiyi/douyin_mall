package service

import (
	"context"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/auth/utils/jwt"
	"douyin_mall/common/constant"
)

type RefreshTokenByRPCService struct {
	ctx context.Context
} // NewRefreshTokenByRPCService new RefreshTokenByRPCService
func NewRefreshTokenByRPCService(ctx context.Context) *RefreshTokenByRPCService {
	return &RefreshTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *RefreshTokenByRPCService) Run(req *auth.RefreshTokenReq) (resp *auth.RefreshTokenResp, err error) {
	newAccessToken, newRefreshToken, success := jwt.RefreshAccessToken(req.RefreshToken)
	if success {
		resp = &auth.RefreshTokenResp{
			StatusCode:   0,
			StatusMsg:    "success",
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}
		return
	}
	resp = &auth.RefreshTokenResp{
		StatusCode: 1006,
		StatusMsg:  constant.GetMsg(1006),
	}
	return
}
