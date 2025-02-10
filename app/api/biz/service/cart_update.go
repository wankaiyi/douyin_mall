package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type CartUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCartUpdateService(Context context.Context, RequestContext *app.RequestContext) *CartUpdateService {
	return &CartUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *CartUpdateService) Run(req *order.CartUpdateRequest) (resp *order.CartUpdateResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
