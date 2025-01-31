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
