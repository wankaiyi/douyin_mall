package service

import (
	"context"
	payment "douyin_mall/payment/kitex_gen/payment"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		Amount:  999,
		OrderId: "1634181837",
		UserId:  123456,
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
