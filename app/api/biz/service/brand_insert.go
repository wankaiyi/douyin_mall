package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type BrandInsertService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBrandInsertService(Context context.Context, RequestContext *app.RequestContext) *BrandInsertService {
	return &BrandInsertService{RequestContext: RequestContext, Context: Context}
}

func (h *BrandInsertService) Run(req *product.BrandInsertRequest) (resp *product.BrandInsertResponse, err error) {
	brand, err := rpc.ProductClient.InsertBrand(h.Context,
		&rpcproduct.BrandInsertReq{
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
		},
	)
	if err != nil {
		resp = &product.BrandInsertResponse{
			StatusCode: 500,
			StatusMsg:  constant.GetMsg(500),
		}
		return
	}
	resp = &product.BrandInsertResponse{
		StatusCode: brand.StatusCode,
		StatusMsg:  brand.StatusMsg,
	}
	return
}
