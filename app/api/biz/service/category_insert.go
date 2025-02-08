package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryInsertService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryInsertService(Context context.Context, RequestContext *app.RequestContext) *CategoryInsertService {
	return &CategoryInsertService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryInsertService) Run(req *product.CategoryInsertRequest) (resp *product.CategoryInsertResponse, err error) {
	category, err := rpc.ProductClient.InsertCategory(h.Context,
		&rpcproduct.CategoryInsertReq{
			Name:        req.Name,
			Description: req.Description,
		},
	)
	if err != nil {
		resp = &product.CategoryInsertResponse{
			StatusCode: 6014,
			StatusMsg:  constant.GetMsg(6014),
		}
		return
	}
	resp = &product.CategoryInsertResponse{
		StatusCode: category.StatusCode,
		StatusMsg:  category.StatusMsg,
	}
	return
}
