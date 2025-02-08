package service

import (
	"context"
	payment "douyin_mall/payment/kitex_gen/payment"
	"testing"
)

func TestNotifyPayment_Run(t *testing.T) {
	ctx := context.Background()
	s := NewNotifyPaymentService(ctx)
	// init req and assert value
	var notifyData = make(map[string]string)
	notifyData["orderId"] = "123456"
	notifyData["tradeStatus"] = "TRADE_CLOSED"

	req := &payment.NotifyPaymentReq{
		NotifyData: notifyData,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
