package service

import (
	"context"
	cart "douyin_mall/cart/kitex_gen/cart"
)

type DeleteCartService struct {
	ctx context.Context
} // NewDeleteCartService new DeleteCartService
func NewDeleteCartService(ctx context.Context) *DeleteCartService {
	return &DeleteCartService{ctx: ctx}
}

// Run create note info
func (s *DeleteCartService) Run(req *cart.DeleteCartReq) (resp *cart.DeleteCartResp, err error) {
	// Finish your business logic.

	return
}
