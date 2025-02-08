package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type InsertCategoryService struct {
	ctx context.Context
} // NewInsertCategoryService new InsertCategoryService
func NewInsertCategoryService(ctx context.Context) *InsertCategoryService {
	return &InsertCategoryService{ctx: ctx}
}

// Run create note info
func (s *InsertCategoryService) Run(req *product.CategoryInsertReq) (resp *product.CategoryInsertResp, err error) {
	// Finish your business logic.
	//将数据封装到结构体中
	category := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	//调用插入数据库的方法
	err = model.CreateCategory(mysql.DB, s.ctx, &category)
	if err != nil {
		klog.Error("insert category failed, error:%v", err)
		resp = &product.CategoryInsertResp{
			StatusCode: 6014,
			StatusMsg:  constant.GetMsg(6014),
		}
		return
	}
	//返回响应
	resp = &product.CategoryInsertResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
