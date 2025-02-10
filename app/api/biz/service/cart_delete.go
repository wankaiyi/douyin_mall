package service

import (
	"context"

	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type CartDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCartDeleteService(Context context.Context, RequestContext *app.RequestContext) *CartDeleteService {
	return &CartDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *CartDeleteService) Run(req *order.CartDeleteRequest) (resp *order.CartDeleteResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
