package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.Error("insert category failed, error:%v", err)
		resp = &product.CategoryDeleteResp{
			StatusCode: 6015,
			StatusMsg:  constant.GetMsg(6015),
		}
		return
	}
	//返回响应
	resp = &product.CategoryDeleteResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
