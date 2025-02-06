package service

import (
	"context"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/auth/utils/jwt"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	ctx := s.ctx
	refreshToken, err := jwt.GenerateRefreshToken(ctx, req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, "生成refresh token失败，req: %v, error: %v", req, err)
		return nil, err
	}
	accessToken, err := jwt.GenerateAccessToken(ctx, req.UserId, req.Role)
	if err != nil {
		klog.CtxErrorf(ctx, "生成access token失败，req: %v, error: %v", req, err)
		return nil, err
	}
	return &auth.DeliveryResp{
		StatusCode:   0,
		StatusMsg:    "success",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
