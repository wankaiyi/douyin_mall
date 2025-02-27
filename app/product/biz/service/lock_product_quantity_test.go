package service

import (
	"context"
	product "douyin_mall/product/kitex_gen/product"
	
)

func TestLockProductQuantity_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLockProductQuantityService(ctx)
	// init req and assert value

	req := &product.ProductLockQuantityRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
