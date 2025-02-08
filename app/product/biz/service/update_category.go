package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
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
		ID:          req.CategoryId,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		resp = &product.CategoryUpdateResp{
			StatusCode: 6016,
			StatusMsg:  constant.GetMsg(6016),
		}
		return
	}
	resp = &product.CategoryUpdateResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
