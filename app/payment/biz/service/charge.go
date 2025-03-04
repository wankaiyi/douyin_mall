package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/payment/biz/dal/alipay"
	kafkaConstant "douyin_mall/payment/infra/kafka/constant"
	"douyin_mall/payment/infra/kafka/producer"
	"time"

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
	ctx := s.ctx
	klog.CtxInfof(ctx, "支付接口被调用, req: %+v", req)
	orderId, err := strconv.ParseInt(req.OrderId, 0, 64)
	if err != nil {
		klog.CtxErrorf(ctx, "parse order id error: %s", err.Error())
		return nil, errors.WithStack(err)
	}
	amount := req.Amount
	paymentUrl := ""
	maxRetryTimes := 3
	for i := 0; i < maxRetryTimes; i++ {
		paymentUrl, err = alipay.Pay(ctx, orderId, amount, req.UserId)
		if err != nil {
			klog.CtxErrorf(ctx, "支付失败, 第%d次重试, req: %+v, 错误信息: %+v", i+1, req, err.Error())
			time.Sleep(500 * time.Millisecond)
		}
	}
	if err != nil {
		klog.CtxErrorf(ctx, "支付失败, 重试%d次后仍失败, req: %+v, 错误信息: %+v", maxRetryTimes, req, err.Error())
		return nil, errors.WithStack(err)
	}

	//给kafka发送延时消息
	producer.SendCheckoutDelayMsg(ctx, strconv.FormatInt(orderId, 10), kafkaConstant.DelayTopic1mLevel)

	resp = &payment.ChargeResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		PaymentUrl: paymentUrl,
	}
	klog.CtxInfof(ctx, "支付接口返回, resp: %+v ,支付接口调用结束。", resp)
	return
}
