package product

import (
	"context"
	product "douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func InsertProduct(ctx context.Context, req *product.InsertProductReq, callOptions ...callopt.Option) (resp *product.InsertProductResp, err error) {
	resp, err = defaultClient.InsertProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "InsertProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SelectProduct(ctx context.Context, req *product.SelectProductReq, callOptions ...callopt.Option) (resp *product.SelectProductResp, err error) {
	resp, err = defaultClient.SelectProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SelectProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
