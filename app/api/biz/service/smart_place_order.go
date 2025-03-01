package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcOrder "douyin_mall/rpc/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type SmartPlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSmartPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *SmartPlaceOrderService {
	return &SmartPlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *SmartPlaceOrderService) Run(req *order.SmartPlaceOrderRequest) (resp *order.SmartPlaceOrderResponse, err error) {
	ctx := h.Context
	userId := ctx.Value(constant.UserId).(int32)
	smartPlaceOrderResp, err := rpc.OrderClient.SmartPlaceOrder(ctx,
		&rpcOrder.SmartPlaceOrderReq{
			UserId:   userId,
			Uuid:     req.Uuid,
			Question: req.Question,
		})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用AI下单失败, req: %v, error: %v", req, err)
		return nil, errors.New("服务器繁忙，请稍后再试")
	}
	return &order.SmartPlaceOrderResponse{
		StatusCode: smartPlaceOrderResp.StatusCode,
		StatusMsg:  smartPlaceOrderResp.StatusMsg,
		Response:   smartPlaceOrderResp.Response,
	}, nil
}
