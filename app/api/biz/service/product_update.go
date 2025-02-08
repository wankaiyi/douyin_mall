package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ProductUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewProductUpdateService(Context context.Context, RequestContext *app.RequestContext) *ProductUpdateService {
	return &ProductUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *ProductUpdateService) Run(req *product.ProductUpdateRequest) (resp *product.ProductUpdateResponse, err error) {
	//根据id查询商品信息
	selectProduct, err := rpc.ProductClient.UpdateProduct(h.Context, &rpcproduct.UpdateProductReq{
		Id:            req.GetId(),
		Name:          req.GetName(),
		Price:         req.GetPrice(),
		Stock:         req.GetStock(),
		Sale:          req.GetSale(),
		PublishStatus: req.GetPublishStatus(),
		Description:   req.GetDescription(),
		Picture:       req.GetPicture(),
		Categories:    req.GetCategories(),
		BrandId:       req.GetBrandId(),
		CategoryId:    req.GetCategoryId(),
	})
	//组装返回数据
	return &product.ProductUpdateResponse{
		StatusCode: selectProduct.StatusCode,
		StatusMsg:  selectProduct.StatusMsg,
	}, err
}
