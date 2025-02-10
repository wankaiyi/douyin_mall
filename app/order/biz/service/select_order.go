package service

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
)

type SelectOrderService struct {
	ctx context.Context
} // NewSelectOrderService new SelectOrderService
func NewSelectOrderService(ctx context.Context) *SelectOrderService {
	return &SelectOrderService{ctx: ctx}
}

// Run create note info
func (s *SelectOrderService) Run(req *order.SelectOrderReq) (resp *order.SelectOrderResp, err error) {
	// Finish your business logic.

	return
}
