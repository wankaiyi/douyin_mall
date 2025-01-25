package service

import (
	"context"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/auth/utils"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	refreshToken, err := utils.GenerateRefreshToken(req.UserId)
	if err != nil {
		return nil, err
	}
	accessToken, err := utils.GenerateAccessToken(req.UserId)
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
