package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/payment/biz/dal/alipay"
	kafkaConstant "douyin_mall/payment/infra/kafka/constant"
	"douyin_mall/payment/infra/kafka/producer"

	payment "douyin_mall/payment/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"strconv"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	orderId, err := strconv.ParseInt(req.OrderId, 0, 64)
	if err != nil {
		klog.CtxErrorf(s.ctx, "parse order id error: %s", err.Error())
		return nil, errors.WithStack(err)
	}
	amount := req.Amount

	paymentUrl, err := alipay.Pay(s.ctx, orderId, amount, req.UserId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "pay error: %s,req: %+v", err.Error(), req)
		resp = &payment.ChargeResp{
			StatusCode: 5000,
			StatusMsg:  constant.GetMsg(5000),
			PaymentUrl: "",
		}
		return nil, errors.WithStack(err)
	}
	//给kafka发送延时消息
	producer.SendCheckoutDelayMsg(s.ctx, strconv.FormatInt(orderId, 10), kafkaConstant.DelayTopic1mLevel)

	resp = &payment.ChargeResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		PaymentUrl: paymentUrl,
	}

	return
}
