package service

import (
	"context"
	product "douyin_mall/api/hertz_gen/api/product"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type SearchService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchService(Context context.Context, RequestContext *app.RequestContext) *SearchService {
	return &SearchService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchService) Run(req *product.ProductRequest) (resp *product.ProductResponse, err error) {
	client := rpc.ProductClient
	res, err := client.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.ProductName})
	if err != nil {
		hlog.Error("商品搜索失败", err)
		return nil, errors.New("商品搜索失败")
	}
	productList := []*product.Product{}
	for i := range res.Results {
		source := res.Results[i]
		productList = append(productList, &product.Product{
			Name:        source.Name,
			Description: source.Description,
		})
	}
	resp = &product.ProductResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
		Products:   productList,
	}
	return resp, err
}
