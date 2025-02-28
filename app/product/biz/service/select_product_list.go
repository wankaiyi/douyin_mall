package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SelectProductListService struct {
	ctx context.Context
} // NewSelectProductListService new SelectProductListService
func NewSelectProductListService(ctx context.Context) *SelectProductListService {
	return &SelectProductListService{ctx: ctx}
}

// Run create note info
func (s *SelectProductListService) Run(req *product.SelectProductListReq) (resp *product.SelectProductListResp, err error) {
	// Finish your business logic.
	// 创建实体类
	products, err := model.SelectProductList(mysql.DB, s.ctx, req.Ids)
	if err != nil {
		klog.CtxErrorf(s.ctx, "查询商品列表失败, error:%v", err)
		return &product.SelectProductListResp{
			StatusCode: 6003,
			StatusMsg:  constant.GetMsg(6003),
		}, nil
	}
	var productList []*product.Product
	for i := range products {
		productList = append(productList, &product.Product{
			Id:           products[i].ProductId,
			Name:         products[i].ProductName,
			Description:  products[i].ProductDescription,
			Picture:      products[i].ProductPicture,
			Price:        products[i].ProductPrice,
			CategoryName: products[i].CategoryName,
			CategoryId:   products[i].CategoryID,
		})
	}
	return &product.SelectProductListResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Products:   productList,
	}, nil
}
