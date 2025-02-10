package service

import (
	"context"
	"douyin_mall/payment/biz/dal/alipay"
	payment "douyin_mall/payment/kitex_gen/payment"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	alipay.Init()
	req := &payment.ChargeReq{
		Amount:  999,
		OrderId: "123456789011",
		UserId:  123456,
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
