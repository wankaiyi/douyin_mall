package service

import (
	"context"
	cart "douyin_mall/cart/kitex_gen/cart"
)

type InsertCartService struct {
	ctx context.Context
} // NewInsertCartService new InsertCartService
func NewInsertCartService(ctx context.Context) *InsertCartService {
	return &InsertCartService{ctx: ctx}
}

// Run create note info
func (s *InsertCartService) Run(req *cart.InsertCartReq) (resp *cart.InsertCartResp, err error) {
	// Finish your business logic.

	return
}
