package service

import (
	"context"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/auth/utils/jwt"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	refreshToken, err := jwt.GenerateRefreshToken(req.UserId, req.Role)
	if err != nil {
		return nil, err
	}
	accessToken, err := jwt.GenerateAccessToken(req.UserId, req.Role)
	if err != nil {
		return nil, err
	}
	return &auth.DeliveryResp{
		StatusCode:   0,
		StatusMsg:    "success",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
