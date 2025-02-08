package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type BrandSelectService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBrandSelectService(Context context.Context, RequestContext *app.RequestContext) *BrandSelectService {
	return &BrandSelectService{RequestContext: RequestContext, Context: Context}
}

func (h *BrandSelectService) Run(req *product.BrandSelectRequest) (resp *product.BrandSelectResponse, err error) {
	brand, err := rpc.ProductClient.SelectBrand(h.Context,
		&rpcproduct.BrandSelectReq{
			BrandId: req.BrandId,
		},
	)
	if err != nil {
		resp = &product.BrandSelectResponse{
			StatusCode: 500,
			StatusMsg:  constant.GetMsg(500),
		}
		return
	}
	resp = &product.BrandSelectResponse{
		StatusCode: brand.StatusCode,
		StatusMsg:  brand.StatusMsg,
		Brand: &product.Brand{
			Id:          req.BrandId,
			Name:        brand.Brand.Name,
			Description: brand.Brand.Name,
			Icon:        brand.Brand.Name,
		},
	}
	return
}
