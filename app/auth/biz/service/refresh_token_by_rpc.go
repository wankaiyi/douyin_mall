package service

import (
	"context"
	"douyin_mall/auth/infra/kafka/producer"
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
	userId, newAccessToken, newRefreshToken, success := jwt.RefreshAccessToken(s.ctx, req.RefreshToken)
	if success {
		resp = &auth.RefreshTokenResp{
			StatusCode:   0,
			StatusMsg:    "success",
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}

		producer.SendUserCacheMessage(userId)

		return
	}
	resp = &auth.RefreshTokenResp{
		StatusCode: 1006,
		StatusMsg:  constant.GetMsg(1006),
	}
	return
}
