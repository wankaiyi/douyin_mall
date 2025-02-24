package service

import (
	"context"
	"douyin_mall/cart/biz/dal/mysql"
	"douyin_mall/cart/biz/model"
	cart "douyin_mall/cart/kitex_gen/cart"
	commonConstant "douyin_mall/common/constant"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	ctx := s.ctx
	err = model.AddCartItem(ctx, mysql.DB, &model.CartItem{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "数据库插入cart item失败, req: %v, err: %v", req, err)
		return nil, err
	}
	return &cart.AddItemResp{StatusCode: 0, StatusMsg: commonConstant.GetMsg(0)}, nil
}
