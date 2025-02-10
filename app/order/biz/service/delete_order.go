package service

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
)

type DeleteOrderService struct {
	ctx context.Context
} // NewDeleteOrderService new DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run create note info
func (s *DeleteOrderService) Run(req *order.DeleteOrderReq) (resp *order.DeleteOrderResp, err error) {
	// Finish your business logic.

	return
}
