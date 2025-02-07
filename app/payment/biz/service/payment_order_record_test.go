package service

import (
	"context"
	"douyin_mall/payment/biz/dal/mysql"
	payment "douyin_mall/payment/kitex_gen/payment"
	"testing"
)

func TestPaymentOrderRecord_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPaymentOrderRecordService(ctx)
	// init req and assert value
	mysql.Init()
	req := &payment.PaymentOrderRecordReq{
		OrderId: "123456789",
		UserId:  123123,
		Amount:  100,
		Status:  1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
