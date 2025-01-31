package main

import (
	"context"
	"douyin_mall/payment/biz/service"
	payment "douyin_mall/payment/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}

// CancelCharge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CancelCharge(ctx context.Context, req *payment.CancelChargeReq) (resp *payment.CancelChargeResp, err error) {
	resp, err = service.NewCancelChargeService(ctx).Run(req)

	return resp, err
}
