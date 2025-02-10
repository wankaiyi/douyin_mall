package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderSelectService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderSelectService(Context context.Context, RequestContext *app.RequestContext) *OrderSelectService {
	return &OrderSelectService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderSelectService) Run(req *order.OrderSelectRequest) (resp *order.OrderSelectResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
