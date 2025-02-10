package service

import (
	"context"
	cart "douyin_mall/cart/kitex_gen/cart"
)

type UpdateCartService struct {
	ctx context.Context
} // NewUpdateCartService new UpdateCartService
func NewUpdateCartService(ctx context.Context) *UpdateCartService {
	return &UpdateCartService{ctx: ctx}
}

// Run create note info
func (s *UpdateCartService) Run(req *cart.UpdateCartReq) (resp *cart.UpdateCartResp, err error) {
	// Finish your business logic.

	return
}
