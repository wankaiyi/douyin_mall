package service

import (
	"context"
	"douyin_mall/common/constant"
	payment "douyin_mall/payment/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
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
	orderId, err := strconv.ParseInt(req.OrderId, 0, 64)
	if err != nil {
		klog.Errorf("parse order id error: %s", err.Error())
		return nil, err
	}
	result, err := CancelPay(s.ctx, orderId)
	if err != nil {
		klog.Errorf("cancel charge error: %s", err.Error())
		return nil, err
	}
	if !result {
		klog.Error("cancel charge failed")

		resp = &payment.CancelChargeResp{
			StatusCode: 5004,
			StatusMsg:  constant.GetMsg(5004),
		}
	}
	resp = &payment.CancelChargeResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
