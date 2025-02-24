package service

import (
	"context"
	"douyin_mall/cart/biz/dal/mysql"
	"douyin_mall/cart/biz/model"
	cart "douyin_mall/cart/kitex_gen/cart"
	"douyin_mall/common/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	ctx := s.ctx
	userId := req.UserId

	err = model.EmptyCart(ctx, mysql.DB, userId)
	if err != nil {
		klog.CtxErrorf(ctx, "数据库操作清空购物车失败, userId: %d, err: %v", userId, err)
		return nil, errors.WithStack(err)
	}
	return &cart.EmptyCartResp{StatusCode: 0, StatusMsg: constant.GetMsg(0)}, nil
}
