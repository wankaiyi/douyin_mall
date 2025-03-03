package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"strconv"

	product "douyin_mall/api/hertz_gen/api/product"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ProductSelectService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewProductSelectService(Context context.Context, RequestContext *app.RequestContext) *ProductSelectService {
	return &ProductSelectService{RequestContext: RequestContext, Context: Context}
}

func (h *ProductSelectService) Run(req *product.ProductSelectRequest) (resp *product.ProductSelectResponse, err error) {
	//根据id查询商品信息
	id, err := strconv.ParseInt(h.RequestContext.Param("id"), 10, 64)
	if err != nil {
		return nil, err
	}
	selectProduct, err := rpc.ProductClient.SelectProduct(h.Context, &rpcproduct.SelectProductReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	//组装返回数据
	return &product.ProductSelectResponse{
		StatusCode: 200,
		StatusMsg:  "success",
		Product: &product.Product{
			Id:            selectProduct.Product.Id,
			Name:          selectProduct.Product.Name,
			Description:   selectProduct.Product.Description,
			Price:         selectProduct.Product.Price,
			Stock:         selectProduct.Product.Stock,
			Sale:          selectProduct.Product.Sale,
			PublishStatus: selectProduct.Product.PublishStatus,
			Picture:       selectProduct.Product.Picture,
			BrandId:       selectProduct.Product.BrandId,
			CategoryId:    selectProduct.Product.CategoryId,
		},
	}, nil
}
