package service

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
)

type InsertOrderService struct {
	ctx context.Context
} // NewInsertOrderService new InsertOrderService
func NewInsertOrderService(ctx context.Context) *InsertOrderService {
	return &InsertOrderService{ctx: ctx}
}

// Run create note info
func (s *InsertOrderService) Run(req *order.InsertOrderReq) (resp *order.InsertOrderResp, err error) {
	// Finish your business logic.

	return
}
