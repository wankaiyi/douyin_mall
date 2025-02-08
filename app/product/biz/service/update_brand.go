package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
)

type UpdateBrandService struct {
	ctx context.Context
} // NewUpdateBrandService new UpdateBrandService
func NewUpdateBrandService(ctx context.Context) *UpdateBrandService {
	return &UpdateBrandService{ctx: ctx}
}

// Run create note info
func (s *UpdateBrandService) Run(req *product.BrandUpdateReq) (resp *product.BrandUpdateResp, err error) {
	// 根据id删除商品
	err = model.UpdateBrand(mysql.DB, s.ctx, &model.Brand{
		ID:          req.BrandId,
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	})
	if err != nil {
		resp = &product.BrandUpdateResp{
			StatusCode: 6020,
			StatusMsg:  constant.GetMsg(6020),
		}
		return
	}
	resp = &product.BrandUpdateResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
