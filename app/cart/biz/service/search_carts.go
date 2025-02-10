package service

import (
	"context"
	cart "douyin_mall/cart/kitex_gen/cart"
)

type SearchCartsService struct {
	ctx context.Context
} // NewSearchCartsService new SearchCartsService
func NewSearchCartsService(ctx context.Context) *SearchCartsService {
	return &SearchCartsService{ctx: ctx}
}

// Run create note info
func (s *SearchCartsService) Run(req *cart.SearchCartsReq) (resp *cart.SearchCartsResp, err error) {
	// Finish your business logic.

	return
}
