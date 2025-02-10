package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type CartSelectService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCartSelectService(Context context.Context, RequestContext *app.RequestContext) *CartSelectService {
	return &CartSelectService{RequestContext: RequestContext, Context: Context}
}

func (h *CartSelectService) Run(req *order.CartSelectRequest) (resp *order.CartSelectResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
