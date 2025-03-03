package service

import (
	"context"
	product "douyin_mall/api/hertz_gen/api/product"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoriesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoriesService(Context context.Context, RequestContext *app.RequestContext) *CategoriesService {
	return &CategoriesService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoriesService) Run(req *product.CategoryRequest) (resp *product.CategoryResponse, err error) {
	categories, err := rpc.ProductClient.GetAllCategories(h.Context, &rpcproduct.CategoryListReq{})
	if err != nil {
		return nil, err
	}
	var categoryList = make([]*product.Category, 0)
	for _, category := range categories.Categories {
		categoryList = append(categoryList, &product.Category{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &product.CategoryResponse{
		Categories: categoryList,
		StatusCode: categories.StatusCode,
		StatusMsg:  categories.StatusMsg,
	}, nil
}
