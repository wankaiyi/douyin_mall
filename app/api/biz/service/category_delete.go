package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryDeleteService(Context context.Context, RequestContext *app.RequestContext) *CategoryDeleteService {
	return &CategoryDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryDeleteService) Run(req *product.CategoryDeleteRequest) (resp *product.CategoryDeleteResponse, err error) {
	category, err := rpc.ProductClient.DeleteCategory(h.Context,
		&rpcproduct.CategoryDeleteReq{
			CategoryId: req.CategoryId,
		},
	)
	if err != nil {
		resp = &product.CategoryDeleteResponse{
			StatusCode: 500,
			StatusMsg:  constant.GetMsg(500),
		}
		return
	}
	resp = &product.CategoryDeleteResponse{
		StatusCode: category.StatusCode,
		StatusMsg:  category.StatusMsg,
	}
	return
}
