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

type DeleteCategoryService struct {
	ctx context.Context
} // NewDeleteCategoryService new DeleteCategoryService
func NewDeleteCategoryService(ctx context.Context) *DeleteCategoryService {
	return &DeleteCategoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteCategoryService) Run(req *product.CategoryDeleteReq) (resp *product.CategoryDeleteResp, err error) {
	//调用插入数据库的方法
	err = model.DeleteCategory(mysql.DB, s.ctx, req.CategoryId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "分类数据库删除失败, error:%v", err)
		return nil, errors.WithStack(err)
	}
	//返回响应
	return &product.CategoryDeleteResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
