package service

import (
	"context"
	"douyin_mall/auth/biz/dal/redis"
	auth "douyin_mall/auth/kitex_gen/auth"
	redisUtils "douyin_mall/auth/utils/redis"
	"douyin_mall/common/constant"
	"github.com/pkg/errors"
)

type RevokeTokenByRPCService struct {
	ctx context.Context
} // NewRevokeTokenByRPCService new RevokeTokenByRPCService
func NewRevokeTokenByRPCService(ctx context.Context) *RevokeTokenByRPCService {
	return &RevokeTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *RevokeTokenByRPCService) Run(req *auth.RevokeTokenReq) (resp *auth.RevokeResp, err error) {
	// Finish your business logic.
	err = redis.RedisClient.Del(s.ctx, redisUtils.GetAccessTokenKey(req.UserId), redisUtils.GetRefreshTokenKey(req.UserId)).Err()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &auth.RevokeResp{StatusCode: 0, StatusMsg: constant.GetMsg(0)}, nil
}
