package service

import (
	"context"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/auth/utils/jwt"
	"douyin_mall/auth/utils/redis"
	"douyin_mall/common/constant"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// 检查 accessToken 是否为空
	if req.AccessToken == "" {
		return &auth.VerifyResp{
			StatusCode: 1004,
			StatusMsg:  constant.GetMsg(1004),
		}, nil
	}
	// 校验access token
	userId, err := jwt.GetUserIdFromToken(req.AccessToken)
	if err != nil {
		return &auth.VerifyResp{
			StatusCode: 1004,
			StatusMsg:  constant.GetMsg(1004),
		}, nil
	}

	// 检查 Redis 中的 access token
	savedAccessToken, err := redis.GetVal(s.ctx, redis.GetAccessTokenKey(userId))
	if err != nil {
		return nil, err
	}

	if savedAccessToken != req.AccessToken {
		return &auth.VerifyResp{
			StatusCode: 1005,
			StatusMsg:  constant.GetMsg(1005),
		}, nil
	}

	return &auth.VerifyResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		UserId:     userId,
	}, nil
}
