package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/payment/biz/dal/alipay"
	payment "douyin_mall/payment/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"strconv"
)

type CancelChargeService struct {
	ctx context.Context
} // NewCancelChargeService new CancelChargeService
func NewCancelChargeService(ctx context.Context) *CancelChargeService {
	return &CancelChargeService{ctx: ctx}
}

// Run create note info
func (s *CancelChargeService) Run(req *payment.CancelChargeReq) (resp *payment.CancelChargeResp, err error) {
	// Finish your business logic.
	klog.CtxInfof(s.ctx, "订单orderId: %s请求取消支付")
	orderId, err := strconv.ParseInt(req.OrderId, 0, 64)
	if err != nil {
		klog.CtxErrorf(s.ctx, "parse order id error: %s", err.Error())
		return nil, err
	}
	result, err := alipay.CancelPay(s.ctx, orderId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "cancel charge error: %s", err.Error())
		return nil, errors.WithStack(err)
	}
	if !result {
		klog.CtxErrorf(s.ctx, "cancel charge failed")

		resp = &payment.CancelChargeResp{
			StatusCode: 5004,
			StatusMsg:  constant.GetMsg(5004),
		}
	}
	klog.CtxInfof(s.ctx, "订单orderId: %s取消支付成功")
	resp = &payment.CancelChargeResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}

	return
}
