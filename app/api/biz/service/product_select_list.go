package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ProductSelectListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewProductSelectListService(Context context.Context, RequestContext *app.RequestContext) *ProductSelectListService {
	return &ProductSelectListService{RequestContext: RequestContext, Context: Context}
}

func (h *ProductSelectListService) Run(req *product.ProductSelectListRequest) (resp *product.ProductSelectListResponse, err error) {
	//根据id查询商品信息
	selectProduct, err := rpc.ProductClient.SelectProductList(h.Context, &rpcproduct.SelectProductListReq{Ids: req.Id})
	if err != nil {
		return nil, err
	}
	//组装返回数据
	var products []*product.Product
	for i := range selectProduct.Product {
		products = append(products, &product.Product{
			Id:            selectProduct.Product[i].Id,
			Name:          selectProduct.Product[i].Name,
			Description:   selectProduct.Product[i].Description,
			Price:         selectProduct.Product[i].Price,
			Stock:         selectProduct.Product[i].Stock,
			Sale:          selectProduct.Product[i].Sale,
			PublishStatus: selectProduct.Product[i].PublishStatus,
			Picture:       selectProduct.Product[i].Picture,
			Categories:    selectProduct.Product[i].Categories,
			BrandId:       selectProduct.Product[i].BrandId,
			CategoryId:    selectProduct.Product[i].CategoryId,
		})
	}

	return &product.ProductSelectListResponse{
		StatusCode: 200,
		StatusMsg:  "success",
		Products:   products,
	}, nil
}
