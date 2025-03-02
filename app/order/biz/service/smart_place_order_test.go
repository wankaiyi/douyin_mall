package service

import (
	"context"
	"douyin_mall/order/biz/dal"
	"douyin_mall/order/infra/rpc"
	order "douyin_mall/order/kitex_gen/order"
	"testing"
)

func TestSmartPlaceOrder_Run(t *testing.T) {
	dal.Init()
	rpc.InitClient()
	ctx := context.Background()
	s := NewSmartPlaceOrderService(ctx)
	// init req and assert value

	req := &order.SmartPlaceOrderReq{
		UserId:   13,
		Uuid:     "123456",
		Question: "你好，请问有什么可以帮到您的？",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
