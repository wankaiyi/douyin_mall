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

// PaymentOrderRecord implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) PaymentOrderRecord(ctx context.Context, req *payment.PaymentOrderRecordReq) (resp *payment.PaymentOrderRecordResp, err error) {
	resp, err = service.NewPaymentOrderRecordService(ctx).Run(req)

	return resp, err
}

// PaymentTransactionRecord implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) PaymentTransactionRecord(ctx context.Context, req *payment.PaymentTransactionRecordReq) (resp *payment.PaymentTransactionRecordResp, err error) {
	resp, err = service.NewPaymentTransactionRecordService(ctx).Run(req)

	return resp, err
}

// IdempotentControl implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) IdempotentControl(ctx context.Context, req *payment.IdempotentControlReq) (resp *payment.IdempotentControlResp, err error) {
	resp, err = service.NewIdempotentControlService(ctx).Run(req)

	return resp, err
}
