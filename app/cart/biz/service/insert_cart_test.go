package service

import (
	"context"
	cart "douyin_mall/cart/kitex_gen/cart"
	"testing"
)

func TestInsertCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewInsertCartService(ctx)
	// init req and assert value

	req := &cart.InsertCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
