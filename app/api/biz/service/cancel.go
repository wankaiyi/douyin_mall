package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcpayment "douyin_mall/rpc/kitex_gen/payment"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	payment "douyin_mall/api/hertz_gen/api/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type CancelService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCancelService(Context context.Context, RequestContext *app.RequestContext) *CancelService {
	return &CancelService{RequestContext: RequestContext, Context: Context}
}

func (h *CancelService) Run(req *payment.CancelRequest) (resp *payment.CancelResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	if req.OrderId == "" {
		return nil, errors.New("订单号不能为空")
	}
	client := rpc.PaymentClient
	res, err := client.CancelCharge(h.Context, &rpcpayment.CancelChargeReq{
		OrderId: req.OrderId,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "cancel charge failed, err: %s, req: %+v", err, &req)
		return nil, errors.New("取消支付失败，请稍后再试")
	}
	resp = &payment.CancelResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}
	return

}
