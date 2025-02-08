package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type BrandUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBrandUpdateService(Context context.Context, RequestContext *app.RequestContext) *BrandUpdateService {
	return &BrandUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *BrandUpdateService) Run(req *product.BrandUpdateRequest) (resp *product.BrandUpdateResponse, err error) {
	brand, err := rpc.ProductClient.UpdateBrand(h.Context,
		&rpcproduct.BrandUpdateReq{
			BrandId:     req.BrandId,
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
		},
	)
	if err != nil {
		resp = &product.BrandUpdateResponse{
			StatusCode: 500,
			StatusMsg:  constant.GetMsg(500),
		}
		return
	}
	resp = &product.BrandUpdateResponse{
		StatusCode: brand.StatusCode,
		StatusMsg:  brand.StatusMsg,
	}
	return
}
