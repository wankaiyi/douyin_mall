package service

import (
	"context"

	product "douyin_mall/api/hertz_gen/api/product"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategorySelectService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategorySelectService(Context context.Context, RequestContext *app.RequestContext) *CategorySelectService {
	return &CategorySelectService{RequestContext: RequestContext, Context: Context}
}

func (h *CategorySelectService) Run(req *product.CategorySelectRequest) (resp *product.CategorySelectResponse, err error) {
	category, err := rpc.ProductClient.SelectCategory(h.Context,
		&rpcproduct.CategorySelectReq{
			CategoryId: req.CategoryId,
		},
	)
	if err != nil {
		resp = &product.CategorySelectResponse{
			StatusCode: category.StatusCode,
			StatusMsg:  category.StatusMsg,
		}
		return
	}
	resp = &product.CategorySelectResponse{
		StatusCode: category.StatusCode,
		StatusMsg:  category.StatusMsg,
		Categories: &product.Category{
			Id:          category.Category.Id,
			Name:        category.Category.Name,
			Description: category.Category.Description,
		},
	}
	return
}
