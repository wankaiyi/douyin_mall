package service

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
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
	result := mysql.DB.Table("tb_product").Delete(&model.Product{ID: req.Id})
	if result.Error != nil {
		err = result.Error
		return &product.DeleteProductResp{StatusCode: 1, StatusMsg: "error"}, result.Error
	}
	return &product.DeleteProductResp{StatusCode: 0, StatusMsg: "success"}, nil
}
