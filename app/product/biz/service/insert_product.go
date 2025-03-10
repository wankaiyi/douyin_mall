package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	producerModel "douyin_mall/product/infra/kafka/model"
	"douyin_mall/product/infra/kafka/producer"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
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
		Name:          req.Name,
		Price:         req.Price,
		Description:   req.Description,
		Stock:         req.Stock,
		Picture:       req.Picture,
		Sale:          0,
		PublishStatus: 1,
		LockStock:     req.Stock,
		CategoryId:    req.CategoryId,
	}
	//调用插入数据库的方法
	err = model.CreateProduct(mysql.DB, s.ctx, &pro)
	if err != nil {
		klog.CtxErrorf(s.ctx, "产品数据库插入失败, error:%v", err)
		return nil, errors.WithStack(err)
	}
	//发送到kafka
	defer func() {
		err := producer.AddToKafka(s.ctx, producerModel.AddProductSendToKafka{
			ID:          pro.ID,
			Name:        pro.Name,
			Price:       pro.Price,
			Description: pro.Description,
			Picture:     pro.Picture,
			Stock:       pro.Stock,
			LockStock:   pro.LockStock,
		})
		if err != nil {
			klog.CtxErrorf(s.ctx, "产品数据库插入信号发送kafka失败, error:%v", err)
		}
	}()
	//返回响应
	return &product.InsertProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
