package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type CartInsertService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCartInsertService(Context context.Context, RequestContext *app.RequestContext) *CartInsertService {
	return &CartInsertService{RequestContext: RequestContext, Context: Context}
}

func (h *CartInsertService) Run(req *order.CartInsertRequest) (resp *order.CartInsertResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
