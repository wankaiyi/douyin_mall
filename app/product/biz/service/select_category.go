package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SelectCategoryService struct {
	ctx context.Context
} // NewSelectCategoryService new SelectCategoryService
func NewSelectCategoryService(ctx context.Context) *SelectCategoryService {
	return &SelectCategoryService{ctx: ctx}
}

// Run create note info
func (s *SelectCategoryService) Run(req *product.CategorySelectReq) (resp *product.CategorySelectResp, err error) {
	category, err := model.SelectCategory(mysql.DB, s.ctx, req.CategoryId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "查询分类失败, err: %v", err)
		return nil, err
	}
	return &product.CategorySelectResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Category: &product.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}
