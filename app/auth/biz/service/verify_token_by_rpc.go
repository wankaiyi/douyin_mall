package service

import (
	"context"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/auth/utils"
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
	// 校验access token
	_, jwtStatus := utils.ParseJWT(req.AccessToken)
	if jwtStatus == utils.TokenValid {
		savedAccessToken, err := redis.GetVal(s.ctx, req.AccessToken)
		if err != nil {
			return nil, err
		}
		if savedAccessToken == req.AccessToken {
			resp = &auth.VerifyResp{
				StatusCode: 0,
				StatusMsg:  constant.GetMsg(0),
				Res:        true,
			}
		} else {
			resp = &auth.VerifyResp{
				StatusCode: 1005,
				StatusMsg:  constant.GetMsg(1005),
				Res:        false,
			}
		}
	} else {
		resp = &auth.VerifyResp{
			StatusCode: 1004,
			StatusMsg:  constant.GetMsg(1004),
			Res:        false,
		}
	}
	return
}
