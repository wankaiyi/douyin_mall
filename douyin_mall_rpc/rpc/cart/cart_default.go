package cart

import (
	"context"
	cart "douyin_mall/rpc/kitex_gen/cart"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AddItem(ctx context.Context, req *cart.AddItemReq, callOptions ...callopt.Option) (resp *cart.AddItemResp, err error) {
	resp, err = defaultClient.AddItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (resp *cart.GetCartResp, err error) {
	resp, err = defaultClient.GetCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (resp *cart.EmptyCartResp, err error) {
	resp, err = defaultClient.EmptyCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "EmptyCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func InsertCart(ctx context.Context, req *cart.InsertCartReq, callOptions ...callopt.Option) (resp *cart.InsertCartResp, err error) {
	resp, err = defaultClient.InsertCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "InsertCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteCart(ctx context.Context, req *cart.DeleteCartReq, callOptions ...callopt.Option) (resp *cart.DeleteCartResp, err error) {
	resp, err = defaultClient.DeleteCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateCart(ctx context.Context, req *cart.UpdateCartReq, callOptions ...callopt.Option) (resp *cart.UpdateCartResp, err error) {
	resp, err = defaultClient.UpdateCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
