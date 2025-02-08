package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type InsertBrandService struct {
	ctx context.Context
} // NewInsertBrandService new InsertBrandService
func NewInsertBrandService(ctx context.Context) *InsertBrandService {
	return &InsertBrandService{ctx: ctx}
}

// Run create note info
func (s *InsertBrandService) Run(req *product.BrandInsertReq) (resp *product.BrandInsertResp, err error) {
	// Finish your business logic.
	//将数据封装到结构体中
	brand := model.Brand{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}
	//调用插入数据库的方法
	err = model.CreateBrand(mysql.DB, s.ctx, &brand)
	if err != nil {
		klog.Error("insert category failed, error:%v", err)
		resp = &product.BrandInsertResp{
			StatusCode: 6018,
			StatusMsg:  constant.GetMsg(6018),
		}
		return
	}
	//返回响应
	resp = &product.BrandInsertResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
