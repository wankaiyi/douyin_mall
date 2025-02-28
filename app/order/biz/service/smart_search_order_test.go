package service

import (
	"context"
	"douyin_mall/order/biz/dal"
	"douyin_mall/order/infra/rpc"
	order "douyin_mall/order/kitex_gen/order"
	"testing"
)

func TestSmartSearchOrder_Run(t *testing.T) {
	dal.Init()
	rpc.InitClient()
	ctx := context.Background()
	s := NewSmartSearchOrderService(ctx)
	// init req and assert value

	req := &order.SmartSearchOrderReq{
		UserId:   13,
		Uuid:     "123456",
		Question: "查一下我买的消毒机",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
