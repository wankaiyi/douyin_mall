package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type DeleteBrandService struct {
	ctx context.Context
} // NewDeleteBrandService new DeleteBrandService
func NewDeleteBrandService(ctx context.Context) *DeleteBrandService {
	return &DeleteBrandService{ctx: ctx}
}

// Run create note info
func (s *DeleteBrandService) Run(req *product.BrandDeleteReq) (resp *product.BrandDeleteResp, err error) {
	//调用插入数据库的方法
	err = model.DeleteBrand(mysql.DB, s.ctx, req.BrandId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "品牌数据库删除失败, error:%v", err)
		return nil, errors.WithStack(err)
	}
	//返回响应
	return &product.BrandDeleteResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
