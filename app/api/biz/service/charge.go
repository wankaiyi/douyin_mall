package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcpayment "douyin_mall/rpc/kitex_gen/payment"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"

	payment "douyin_mall/api/hertz_gen/api/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type ChargeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChargeService(Context context.Context, RequestContext *app.RequestContext) *ChargeService {
	return &ChargeService{RequestContext: RequestContext, Context: Context}
}

func (h *ChargeService) Run(req *payment.ChargeRequest) (resp *payment.ChargeResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	client := rpc.PaymentClient
	res, err := client.Charge(h.Context, &rpcpayment.ChargeReq{
		Amount:  req.Amount,
		OrderId: req.OrderId,
		UserId:  req.UserId,
	})
	if err != nil {
		klog.Error("payment failed, err: ", err)
		return nil, errors.New("支付失败，请稍后再试")
	}
	resp = &payment.ChargeResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}
	return
}
