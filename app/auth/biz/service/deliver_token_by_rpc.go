package service

import (
	"context"
	"douyin_mall/auth/infra/kafka/producer"
	auth "douyin_mall/auth/kitex_gen/auth"
	"douyin_mall/auth/utils/jwt"
	"douyin_mall/common/constant"
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
	userId := req.UserId
	refreshToken, err := jwt.GenerateRefreshToken(ctx, userId)
	if err != nil {
		klog.CtxErrorf(ctx, "生成refresh token失败，req: %v, error: %v", req, err)
		return nil, err
	}
	accessToken, err := jwt.GenerateAccessToken(ctx, userId, req.Role)
	if err != nil {
		klog.CtxErrorf(ctx, "生成access token失败，req: %v, error: %v", req, err)
		return nil, err
	}

	producer.SendUserCacheMessage(ctx, userId)

	return &auth.DeliveryResp{
		StatusCode:   0,
		StatusMsg:    constant.GetMsg(0),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
