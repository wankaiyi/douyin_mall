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
		klog.Error("mysql error:%v", err)
		resp = &product.SelectProductListResp{
			StatusCode: 6003,
			StatusMsg:  constant.GetMsg(6003),
		}
		return
	}
	var productList []*product.Product
	for i := range products {
		productList = append(productList, &product.Product{
			Id:          products[i].ID,
			Name:        products[i].Name,
			Description: products[i].Description,
			Picture:     products[i].Picture,
			Price:       products[i].Price,
		})
	}
	resp = &product.SelectProductListResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Product:    productList,
	}
	return
}
