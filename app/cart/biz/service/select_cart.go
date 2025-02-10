package service

import (
	"context"
	cart "douyin_mall/cart/kitex_gen/cart"
)

type SelectCartService struct {
	ctx context.Context
} // NewSelectCartService new SelectCartService
func NewSelectCartService(ctx context.Context) *SelectCartService {
	return &SelectCartService{ctx: ctx}
}

// Run create note info
func (s *SelectCartService) Run(req *cart.SelectCartReq) (resp *cart.SelectCartResp, err error) {
	// Finish your business logic.

	return
}
