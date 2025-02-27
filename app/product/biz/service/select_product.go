package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SelectProductService struct {
	ctx context.Context
} // NewSelectProductService new SelectProductService
func NewSelectProductService(ctx context.Context) *SelectProductService {
	return &SelectProductService{ctx: ctx}
}

// Run create note info
func (s *SelectProductService) Run(req *product.SelectProductReq) (resp *product.SelectProductResp, err error) {
	pro, err := model.SelectProduct(mysql.DB, s.ctx, req.Id)
	if err != nil {
		klog.CtxErrorf(s.ctx, "查询商品失败, error:%v", err)
		return nil, err
	}
	return &product.SelectProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Product: &product.Product{
			Id:            pro.ID,
			Name:          pro.Name,
			Description:   pro.Description,
			Picture:       pro.Picture,
			Price:         pro.Price,
			Stock:         pro.Stock,
			Sale:          pro.Sale,
			PublishStatus: pro.PublicState,
			CategoryId:    pro.CategoryId,
			BrandId:       pro.BrandId,
		},
	}, nil
}
