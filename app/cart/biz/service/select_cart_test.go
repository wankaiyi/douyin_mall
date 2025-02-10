package service

import (
	"context"
	cart "douyin_mall/cart/kitex_gen/cart"
	"testing"
)

func TestSelectCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSelectCartService(ctx)
	// init req and assert value

	req := &cart.SelectCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
