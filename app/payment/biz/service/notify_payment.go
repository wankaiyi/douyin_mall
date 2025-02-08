package service

import (
	"context"
	"douyin_mall/common/constant"
	payment "douyin_mall/payment/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
)

type NotifyPaymentService struct {
	ctx context.Context
} // NewNotifyPaymentService new NotifyPaymentService
func NewNotifyPaymentService(ctx context.Context) *NotifyPaymentService {
	return &NotifyPaymentService{ctx: ctx}
}

// Run create note info
func (s *NotifyPaymentService) Run(req *payment.NotifyPaymentReq) (resp *payment.NotifyPaymentResp, err error) {
	// Finish your business logic.
	notifyData := req.NotifyData
	//alipayTradeNo := notifyData["alipayTradeNo"]
	tradeStatus := notifyData["tradeStatus"]
	//alipayAmount := notifyData["alipayAmount"]
	orderId := notifyData["orderId"]
	//todo:通过订单号查询订单记录，检查订单状态是否已支付，未支付则更新订单状态，已支付则直接返回
	//todo:通过订单得到userId,得到订单金额
	if tradeStatus != "TRADE_SUCCESS" {
		klog.CtxErrorf(s.ctx, "orderId:%s,pay filed,tradeStatus:%s", orderId, tradeStatus)
		resp = &payment.NotifyPaymentResp{
			StatusCode: 5006,
			StatusMsg:  constant.GetMsg(5006),
		}
		return resp, nil
	}

	resp = &payment.NotifyPaymentResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}

	return resp, nil
}
