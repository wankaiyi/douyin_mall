package category

import (
	"context"
	product "douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SelectCategory(ctx context.Context, req *product.CategorySelectReq, callOptions ...callopt.Option) (resp *product.CategorySelectResp, err error) {
	resp, err = defaultClient.SelectCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SelectCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
