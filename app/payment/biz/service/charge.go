package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/payment/biz/dal/alipay"
	"douyin_mall/payment/infra/kafka"
	payment "douyin_mall/payment/kitex_gen/payment"
	"github.com/IBM/sarama"
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

	paymentUrl, err := alipay.Pay(s.ctx, orderId, amount)
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
	producer := kafka.GetProducer()
	msg := strconv.Itoa(int(orderId))

	kafka.SendDelayMsg(&sarama.ProducerMessage{
		// todo 防止掉单主动延时查询支付宝，5秒太短了，可以设置一个时间梯度，比如10秒，30秒，1分钟
		Topic: "__delay-seconds-5",
		Value: sarama.StringEncoder(msg),
		Key:   sarama.StringEncoder("check-payment"),
	}, *producer)
	resp = &payment.ChargeResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		PaymentUrl: paymentUrl,
	}

	return
}
