package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderDeleteService(Context context.Context, RequestContext *app.RequestContext) *OrderDeleteService {
	return &OrderDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderDeleteService) Run(req *order.OrderDeleteRequest) (resp *order.OrderDeleteResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
