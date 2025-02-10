package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderUpdateService(Context context.Context, RequestContext *app.RequestContext) *OrderUpdateService {
	return &OrderUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderUpdateService) Run(req *order.OrderUpdateRequest) (resp *order.OrderUpdateResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
