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
	var pro model.ProductWithCategory
	err = getProductFromCache(req.Id, &pro)
	if err != nil {
		klog.CtxInfof(s.ctx, "从缓存查询商品失败, error:%v", err)
	}
	list, err := model.SelectProductList(mysql.DB, s.ctx, []int32{req.Id})
	if err != nil || len(list) <= 0 {
		klog.CtxErrorf(s.ctx, "查询商品失败, error:%v", err)
		return nil, err
	}
	klog.CtxInfof(s.ctx, "查询商品成功, product:%v", list)

	pro = list[0]
	return &product.SelectProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Product: &product.Product{
			Id:            pro.ProductId,
			Name:          pro.ProductName,
			Description:   pro.ProductDescription,
			Picture:       pro.ProductPicture,
			Price:         pro.ProductPrice,
			Stock:         pro.ProductStock,
			Sale:          pro.ProductSale,
			PublishStatus: pro.ProductPublicState,
			CategoryId:    pro.CategoryID,
			CategoryName:  pro.CategoryName,
		},
	}, nil
}

func getProductFromCache(id int32, pro *model.ProductWithCategory) (err error) {
	return nil
}
