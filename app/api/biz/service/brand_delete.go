package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type BrandDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBrandDeleteService(Context context.Context, RequestContext *app.RequestContext) *BrandDeleteService {
	return &BrandDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *BrandDeleteService) Run(req *product.BrandDeleteRequest) (resp *product.BrandDeleteResponse, err error) {
	category, err := rpc.ProductClient.DeleteBrand(h.Context,
		&rpcproduct.BrandDeleteReq{
			BrandId: req.BrandId,
		},
	)
	if err != nil {
		resp = &product.BrandDeleteResponse{
			StatusCode: 500,
			StatusMsg:  constant.GetMsg(500),
		}
		return
	}
	resp = &product.BrandDeleteResponse{
		StatusCode: category.StatusCode,
		StatusMsg:  category.StatusMsg,
	}
	return
}
