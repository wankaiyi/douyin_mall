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
	for i := range selectProduct.Products {
		products = append(products, &product.Product{
			Id:            selectProduct.Products[i].Id,
			Name:          selectProduct.Products[i].Name,
			Description:   selectProduct.Products[i].Description,
			Price:         selectProduct.Products[i].Price,
			Stock:         selectProduct.Products[i].Stock,
			Sale:          selectProduct.Products[i].Sale,
			PublishStatus: selectProduct.Products[i].PublishStatus,
			Picture:       selectProduct.Products[i].Picture,
			BrandId:       selectProduct.Products[i].BrandId,
			CategoryId:    selectProduct.Products[i].CategoryId,
			CategoryName:  selectProduct.Products[i].CategoryName,
		})
	}

	return &product.ProductSelectListResponse{
		StatusCode: 200,
		StatusMsg:  "success",
		Products:   products,
	}, nil
}
