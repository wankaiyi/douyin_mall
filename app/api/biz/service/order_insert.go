package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderInsertService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderInsertService(Context context.Context, RequestContext *app.RequestContext) *OrderInsertService {
	return &OrderInsertService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderInsertService) Run(req *order.OrderInsertRequest) (resp *order.OrderInsertResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
