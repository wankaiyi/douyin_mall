package service

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type InsertProductService struct {
	ctx context.Context
} // NewInsertProductService new InsertProductService
func NewInsertProductService(ctx context.Context) *InsertProductService {
	return &InsertProductService{ctx: ctx}
}

// Run create note info
func (s *InsertProductService) Run(req *product.InsertProductReq) (resp *product.InsertProductResp, err error) {
	//将数据封装到结构体中
	pro := model.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Stock:       req.Stock,
		Picture:     req.Picture,
		Sale:        0,
		PublicState: 1,
		LockStock:   req.Stock,
	}
	//调用插入数据库的方法
	result := mysql.DB.Table("tb_product").Create(&pro)
	if result.Error != nil {
		hlog.Error("insert product error:%v", result.Error)
		return nil, result.Error
	}
	//返回响应
	return &product.InsertProductResp{StatusCode: 0, StatusMsg: "success"}, nil
}
