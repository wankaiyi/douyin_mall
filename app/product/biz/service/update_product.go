package service

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
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
	result := mysql.DB.Table("tb_product").Updates(&model.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Stock:       req.Stock,
		Sale:        req.Sale,
		PublicState: req.PublishStatus,
		LockStock:   req.Stock,
	})
	if result.Error != nil {
		err = result.Error
		resp = &product.UpdateProductResp{StatusCode: 1, StatusMsg: "update product failed"}
		return
	}
	if result.RowsAffected == 0 {
		resp = &product.UpdateProductResp{StatusCode: 0, StatusMsg: "not found product"}
		return
	}
	resp = &product.UpdateProductResp{StatusCode: 0, StatusMsg: "success"}
	return
}
