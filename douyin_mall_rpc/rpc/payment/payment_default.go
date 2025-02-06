package payment

import (
	"context"
	payment "douyin_mall/rpc/kitex_gen/payment"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (resp *payment.ChargeResp, err error) {
	resp, err = defaultClient.Charge(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Charge call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CancelCharge(ctx context.Context, req *payment.CancelChargeReq, callOptions ...callopt.Option) (resp *payment.CancelChargeResp, err error) {
	resp, err = defaultClient.CancelCharge(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CancelCharge call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PaymentOrderRecord(ctx context.Context, req *payment.PaymentOrderRecordReq, callOptions ...callopt.Option) (resp *payment.PaymentOrderRecordResp, err error) {
	resp, err = defaultClient.PaymentOrderRecord(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PaymentOrderRecord call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PaymentTransactionRecord(ctx context.Context, req *payment.PaymentTransactionRecordReq, callOptions ...callopt.Option) (resp *payment.PaymentTransactionRecordResp, err error) {
	resp, err = defaultClient.PaymentTransactionRecord(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PaymentTransactionRecord call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func IdempotentControl(ctx context.Context, req *payment.IdempotentControlReq, callOptions ...callopt.Option) (resp *payment.IdempotentControlResp, err error) {
	resp, err = defaultClient.IdempotentControl(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "IdempotentControl call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
