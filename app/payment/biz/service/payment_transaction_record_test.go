package service

import (
	"context"
	"douyin_mall/payment/biz/dal/mysql"
	payment "douyin_mall/payment/kitex_gen/payment"
	"testing"
)

func TestPaymentTransactionRecord_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPaymentTransactionRecordService(ctx)
	// init req and assert value
	mysql.Init()
	req := &payment.PaymentTransactionRecordReq{

		OrderId:       "1234567890",
		AlipayTradeNo: "20210001111122223333444455556666",
		TradeStatus:   "TRADE_SUCCESS",
		Callback:      "http://localhost:8080/payment/callback",
		RequestParams: "request_params",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
