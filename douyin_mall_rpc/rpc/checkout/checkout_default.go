package checkout

import (
	"context"
	checkout "douyin_mall/rpc/kitex_gen/checkout"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (resp *checkout.CheckoutResp, err error) {
	resp, err = defaultClient.Checkout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Checkout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckoutProductItems(ctx context.Context, req *checkout.CheckoutProductItemsReq, callOptions ...callopt.Option) (resp *checkout.CheckoutProductItemsResp, err error) {
	resp, err = defaultClient.CheckoutProductItems(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckoutProductItems call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
