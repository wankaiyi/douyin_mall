package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	kf "douyin_mall/product/infra/kafka"
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
	// 根据id删除商品
	pro := model.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Stock:       req.Stock,
		Sale:        req.Sale,
		PublicState: req.PublishStatus,
		LockStock:   req.Stock,
		CategoryId:  req.CategoryId,
		BrandId:     req.BrandId,
	}
	err = model.UpdateProduct(mysql.DB, s.ctx, &pro)
	if err != nil {
		resp = &product.UpdateProductResp{
			StatusCode: 6001,
			StatusMsg:  constant.GetMsg(6001),
		}
		return
	}
	//发送到kafka
	defer func() {
		err := kf.UpdateProduct(&pro)
		if err != nil {
			klog.Error("update product error:%v", err)
		}
	}()
	resp = &product.UpdateProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(6001),
	}
	return
}
