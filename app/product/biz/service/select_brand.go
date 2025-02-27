package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SelectBrandService struct {
	ctx context.Context
} // NewSelectBrandService new SelectBrandService
func NewSelectBrandService(ctx context.Context) *SelectBrandService {
	return &SelectBrandService{ctx: ctx}
}

// Run create note info
func (s *SelectBrandService) Run(req *product.BrandSelectReq) (resp *product.BrandSelectResp, err error) {
	brand, err := model.SelectBrand(mysql.DB, s.ctx, req.BrandId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "查询品牌失败, error: %v", err)
		return nil, err
	}
	return &product.BrandSelectResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Brand: &product.Brand{
			Id:          brand.ID,
			Name:        brand.Name,
			Description: brand.Description,
		},
	}, nil
}
