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

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// 根据id删除商品
	err = model.DeleteProduct(mysql.DB, s.ctx, req.Id)
	if err != nil {
		klog.CtxErrorf(s.ctx, "产品数据库删除失败, error:%v", err)
		return nil, errors.WithStack(err)
	}
	//发送到kafka
	defer func() {
		err := producer.DeleteToKafka(s.ctx, producerModel.DeleteProductSendToKafka{
			ID: req.Id,
		})
		if err != nil {
			klog.CtxErrorf(s.ctx, "删除信号发送kafka失败, error:%v", err)
		}
	}()
	return &product.DeleteProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
