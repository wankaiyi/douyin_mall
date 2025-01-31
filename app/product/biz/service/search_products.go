package service

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
	query := req.Query
	var p []model.Product
	result := mysql.DB.Table("tb_product").Select("*").Where("name like ?", "%"+query+"%").Find(&p)
	hlog.Error(result)
	resp = &product.SearchProductsResp{
		StatusCode: 0,
		StatusMsg:  "123",
	}
	err = result.Error
	return
}
