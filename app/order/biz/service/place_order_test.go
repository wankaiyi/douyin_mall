package service

import (
	"context"
	"douyin_mall/order/biz/dal"
	"douyin_mall/order/infra/kafka"
	"douyin_mall/order/kitex_gen/cart"
	order "douyin_mall/order/kitex_gen/order"
	"douyin_mall/order/utils"
	"testing"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	dal.Init()
	utils.InitSnowflake()
	kafka.Init()

	s := NewPlaceOrderService(ctx)
	// init req and assert value

	req := &order.PlaceOrderReq{
		UserId: 13,
		Address: &order.Address{
			Name:          "test",
			PhoneNumber:   "12345678901",
			Province:      "Beijing",
			City:          "Beijing",
			Region:        "Beijing",
			DetailAddress: "test",
		},
		OrderItems: []*order.OrderItem{
			{
				Item: &cart.CartItem{
					ProductId: 1,
					Quantity:  1},
				Cost: 99.99,
			},
			{
				Item: &cart.CartItem{
					ProductId: 2,
					Quantity:  2},
				Cost: 199.99,
			},
		},
		TotalCost: 399.98,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
