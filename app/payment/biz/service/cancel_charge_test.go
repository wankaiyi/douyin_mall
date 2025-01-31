package service

import (
	"context"
	payment "douyin_mall/payment/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
)

func TestCancelCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelChargeService(ctx)
	// init req and assert value

	req := &payment.CancelChargeReq{
		OrderId: "1634156123559",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
	klog.Infof("resp: %v", resp)

}
