package service

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
	"testing"
)

func TestSearchOrders_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchOrdersService(ctx)
	// init req and assert value

	req := &order.SearchOrdersReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
