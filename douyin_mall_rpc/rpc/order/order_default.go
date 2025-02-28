package order

import (
	"context"
	order "douyin_mall/rpc/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (resp *order.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = defaultClient.MarkOrderPaid(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MarkOrderPaid call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrder(ctx context.Context, req *order.GetOrderReq, callOptions ...callopt.Option) (resp *order.GetOrderResp, err error) {
	resp, err = defaultClient.GetOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SmartSearchOrder(ctx context.Context, req *order.SmartSearchOrderReq, callOptions ...callopt.Option) (resp *order.SmartSearchOrderResp, err error) {
	resp, err = defaultClient.SmartSearchOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SmartSearchOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
