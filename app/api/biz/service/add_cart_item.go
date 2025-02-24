package service

import (
	"context"
	cart "douyin_mall/api/hertz_gen/api/cart"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcCart "douyin_mall/rpc/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *cart.AddCartItemResp, err error) {
	ctx := h.Context
	addItemResp, err := rpc.CartClient.AddItem(ctx, &rpcCart.AddItemReq{
		Item: &rpcCart.CartItem{
			ProductId: req.Item.ProductId,
			Quantity:  req.Item.Quantity,
		},
		UserId: ctx.Value(constant.UserId).(int32),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用添加购物车失败，req：%v, err: %v", req, err)
		return nil, errors.New("添加购物车失败")
	}
	return &cart.AddCartItemResp{
		StatusCode: addItemResp.StatusCode,
		StatusMsg:  addItemResp.StatusMsg,
	}, nil
}
