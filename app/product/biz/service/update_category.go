package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateCategoryService struct {
	ctx context.Context
} // NewUpdateCategoryService new UpdateCategoryService
func NewUpdateCategoryService(ctx context.Context) *UpdateCategoryService {
	return &UpdateCategoryService{ctx: ctx}
}

// Run create note info
func (s *UpdateCategoryService) Run(req *product.CategoryUpdateReq) (resp *product.CategoryUpdateResp, err error) {
	// 根据id删除商品
	err = model.UpdateCategory(mysql.DB, s.ctx, &model.Category{
		Base: model.Base{
			ID: req.CategoryId,
		},
		Name: req.Name,
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "更新分类失败, err: %v", err)
		return &product.CategoryUpdateResp{
			StatusCode: 6016,
			StatusMsg:  constant.GetMsg(6016),
		}, nil
	}
	return &product.CategoryUpdateResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
