package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	producerModel "douyin_mall/product/infra/kafka/model"
	"douyin_mall/product/infra/kafka/producer"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	pro := model.Product{
		Base: model.Base{
			ID: req.Id,
		},
		Name:          req.Name,
		Description:   req.Description,
		Picture:       req.Picture,
		Price:         req.Price,
		Stock:         req.Stock,
		Sale:          req.Sale,
		PublishStatus: req.PublishStatus,
		LockStock:     req.Stock,
		CategoryId:    req.CategoryId,
	}
	err = model.UpdateProduct(mysql.DB, s.ctx, &pro)
	if err != nil {
		klog.CtxErrorf(s.ctx, "更新商品失败,error:%v", err)
		return nil, err
	}
	//4 调用redis的set方法将数据导入到redis缓存中
	err = model.PushToRedisBaseInfo(s.ctx, model.Product{
		Base: model.Base{
			ID: pro.ID,
		},
		Name:        pro.Name,
		Description: pro.Description,
		Price:       pro.Price,
		Stock:       pro.Stock,
		LockStock:   pro.LockStock,
	}, redis.RedisClient, model.BaseInfoKey(s.ctx, pro.ID))
	if err != nil {
		klog.CtxErrorf(s.ctx, "redis push product to redis err:%v", err)
		return nil, err
	}
	//发送到kafka
	defer func() {
		err := producer.UpdateToKafka(s.ctx, producerModel.UpdateProductSendToKafka{
			ID:          pro.ID,
			Name:        pro.Name,
			Description: pro.Description,
			Price:       pro.Price,
			Picture:     pro.Picture,
			Stock:       pro.Stock,
			LockStock:   pro.LockStock,
		})
		if err != nil {
			klog.CtxErrorf(s.ctx, "推送到kafka失败,error:%v", err)
		}
	}()
	return &product.UpdateProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(6001),
	}, nil
}
