package service

import (
	"context"
	"douyin_mall/payment/biz/dal/mysql"
	payment "douyin_mall/payment/kitex_gen/payment"

	"testing"
)

func TestIdempotentControl_Run(t *testing.T) {
	ctx := context.Background()
	s := NewIdempotentControlService(ctx)
	// init req and assert value
	mysql.Init()
	req := &payment.IdempotentControlReq{
		OrderId: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
