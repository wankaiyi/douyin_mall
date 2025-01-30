package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchService(Context context.Context, RequestContext *app.RequestContext) *SearchService {
	return &SearchService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchService) Run(req *product.ProductRequest) (resp *product.ProductResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	client := rpc.ProductClient
	//把前端传入的参数放进去
	res, err := client.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.ProductName})
	if err != nil {
		klog.Error("payment failed, err: ", err)
		return nil, errors.New("支付失败，请稍后再试")
	}
	resp = &product.ProductResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}
	return resp, err
}
