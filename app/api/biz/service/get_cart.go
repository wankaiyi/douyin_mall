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

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	ctx := h.Context
	getCartResp, err := rpc.CartClient.GetCart(ctx, &rpcCart.GetCartReq{UserId: ctx.Value(constant.UserId).(int32)})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用获取购物车失败，req：%v, err: %v", req, err)
		return nil, errors.New("获取购物车失败")
	}

	items := make([]*cart.Product, len(getCartResp.Products))
	for i, item := range getCartResp.Products {
		items[i] = &cart.Product{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			Price:       item.Price,
			Quantity:    item.Quantity,
		}
	}
	return &cart.GetCartResp{
		StatusCode: getCartResp.StatusCode,
		StatusMsg:  getCartResp.StatusMsg,
		Products:   items,
	}, nil
}
