package service

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
)

type SearchOrdersService struct {
	ctx context.Context
} // NewSearchOrdersService new SearchOrdersService
func NewSearchOrdersService(ctx context.Context) *SearchOrdersService {
	return &SearchOrdersService{ctx: ctx}
}

// Run create note info
func (s *SearchOrdersService) Run(req *order.SearchOrdersReq) (resp *order.SearchOrdersResp, err error) {
	// Finish your business logic.

	return
}
