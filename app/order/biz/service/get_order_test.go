package service

import (
	"context"
	"douyin_mall/order/biz/dal"
	"douyin_mall/order/infra/rpc"
	order "douyin_mall/order/kitex_gen/order"
	"testing"
)

func TestGetOrder_Run(t *testing.T) {
	ctx := context.Background()
	dal.Init()
	rpc.InitClient()
	s := NewGetOrderService(ctx)
	// init req and assert value

	req := &order.GetOrderReq{
		UserId:  13,
		OrderId: "1894961305492680704",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
