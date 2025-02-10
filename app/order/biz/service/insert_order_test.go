package service

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
	"testing"
)

func TestInsertOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewInsertOrderService(ctx)
	// init req and assert value

	req := &order.InsertOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
