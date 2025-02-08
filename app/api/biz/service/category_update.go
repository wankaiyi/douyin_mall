package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryUpdateService(Context context.Context, RequestContext *app.RequestContext) *CategoryUpdateService {
	return &CategoryUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryUpdateService) Run(req *product.CategoryUpdateRequest) (resp *product.CategoryUpdateResponse, err error) {
	category, err := rpc.ProductClient.UpdateCategory(h.Context,
		&rpcproduct.CategoryUpdateReq{
			CategoryId:  req.CategoryId,
			Name:        req.Name,
			Description: req.Description,
		},
	)
	if err != nil {
		resp = &product.CategoryUpdateResponse{
			StatusCode: 500,
			StatusMsg:  constant.GetMsg(500),
		}
		return
	}
	resp = &product.CategoryUpdateResponse{
		StatusCode: category.StatusCode,
		StatusMsg:  category.StatusMsg,
	}
	return
}
