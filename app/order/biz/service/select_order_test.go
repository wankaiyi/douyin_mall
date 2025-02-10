package service

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
	"testing"
)

func TestSelectOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSelectOrderService(ctx)
	// init req and assert value

	req := &order.SelectOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
