package service

import (
	"context"
	"douyin_mall/common/constant"
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
	err = model.DeleteProduct(mysql.DB, s.ctx, req.Id)
	if err != nil {
		resp = &product.DeleteProductResp{
			StatusCode: 6002,
			StatusMsg:  constant.GetMsg(6002),
		}
		return
	}
	resp = &product.DeleteProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
