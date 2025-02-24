package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/task"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
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
		CategoryId:  req.CategoryId,
		BrandId:     req.BrandId,
	}
	//调用插入数据库的方法
	err = model.CreateProduct(mysql.DB, s.ctx, &pro)
	if err != nil {
		klog.Error("insert product error:%v", err)
		resp = &product.InsertProductResp{
			StatusCode: 6000,
			StatusMsg:  constant.GetMsg(6000),
		}
		return
	}
	//发送到kafka
	func() {
		err := task.AddProduct(&pro)
		if err != nil {
			klog.Error("insert product error:%v", err)
		}
	}()
	//返回响应
	resp = &product.InsertProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
