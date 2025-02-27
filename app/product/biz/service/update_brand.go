package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.CtxErrorf(s.ctx, "更新品牌失败, err: %v", err)
		return nil, err
	}
	return &product.BrandUpdateResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
