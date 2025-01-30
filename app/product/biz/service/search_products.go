package service

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/go-pay/xlog"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	db := mysql.DB
	var p []model.Product
	result := db.Table("tb_product").Select("*").Find(&p)
	xlog.Error(result)
	resp = &product.SearchProductsResp{
		StatusCode: 0,
		StatusMsg:  "123",
	}
	err = result.Error
	return
}
