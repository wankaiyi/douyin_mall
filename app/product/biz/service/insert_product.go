package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	kf "douyin_mall/product/infra/kafka"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/segmentio/kafka-go"
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
	//TODO 发送到kafka
	defer sendToKafka(s.ctx, &pro)
	//返回响应
	resp = &product.InsertProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}

func sendToKafka(ctx context.Context, pro *model.Product) {
	productData, err := sonic.Marshal(pro)
	if err != nil {
		hlog.CtxErrorf(ctx, "product序列化失败:%v error:%v", pro, err)
	}
	kfErr := kf.Writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("insert"),
			Value: productData,
		})
	if kfErr != nil {
		klog.CtxErrorf(ctx, "发送kafka失败:%v error:%v", pro, kfErr)
	}
}
