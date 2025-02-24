package service

import (
	"context"
	"douyin_mall/cart/biz/dal/mysql"
	"douyin_mall/cart/biz/model"
	"douyin_mall/cart/infra/rpc"
	cart "douyin_mall/cart/kitex_gen/cart"
	commonConstant "douyin_mall/common/constant"
	"douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	ctx := s.ctx
	cartItems, err := model.GetCartItemByUserId(ctx, mysql.DB, req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, "数据库查询购物车信息失败，req: %v, err: %v", req, err)
		return nil, err
	}

	if len(cartItems) == 0 {
		return &cart.GetCartResp{
			StatusCode: 0,
			StatusMsg:  commonConstant.GetMsg(0),
			Products:   make([]*cart.Product, 0),
		}, nil
	}

	productIds := make([]int64, len(cartItems))
	for i, item := range cartItems {
		productIds[i] = (int64)(item.ProductId)
	}

	productList, err := rpc.ProductClient.SelectProductList(ctx, &product.SelectProductListReq{
		Ids: productIds,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "rpc查询商品信息失败，req: %v, err: %v", req, err)
		return nil, errors.WithStack(err)
	}

	productItems := make([]*cart.Product, len(cartItems))
	for i, item := range cartItems {

		productItems[i] = &cart.Product{
			Id:          item.ProductId,
			Name:        productList.Product[i].Name,
			Description: productList.Product[i].Description,
			Picture:     productList.Product[i].Picture,
			Price:       productList.Product[i].Price,
			Quantity:    item.Quantity,
		}
	}

	return &cart.GetCartResp{
		StatusCode: 0,
		StatusMsg:  commonConstant.GetMsg(0),
		Products:   productItems,
	}, nil
}
