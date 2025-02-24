package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcCart "douyin_mall/rpc/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"

	cart "douyin_mall/api/hertz_gen/api/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	ctx := h.Context
	emptyCartResp, err := rpc.CartClient.EmptyCart(ctx, &rpcCart.EmptyCartReq{
		UserId: ctx.Value(constant.UserId).(int32),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用清空购物车失败, err: %v", err)
		return nil, errors.New("清空购物车失败")
	}
	return &cart.EmptyCartResp{
		StatusCode: emptyCartResp.StatusCode,
		StatusMsg:  emptyCartResp.StatusMsg,
	}, nil
}
