package service

import (
	"context"
	"douyin_mall/api/hertz_gen/api/order"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcOrder "douyin_mall/rpc/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ListOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrderService(Context context.Context, RequestContext *app.RequestContext) *ListOrderService {
	return &ListOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *ListOrderService) Run(req *order.ListOrderRequest) (resp *order.ListOrderResponse, err error) {
	ctx := h.Context
	listOrderResp, err := rpc.OrderClient.ListOrder(ctx, &rpcOrder.ListOrderReq{
		UserId: ctx.Value(constant.UserId).(int32),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用获取订单列表失败, err: %v", err)
		return nil, err
	}
	var orders []*order.Order
	for _, o := range listOrderResp.Orders {
		var products []*order.Product
		for _, p := range o.Products {
			products = append(products, &order.Product{
				Id:          p.Id,
				Name:        p.Name,
				Description: p.Description,
				Picture:     p.Picture,
				Price:       p.Price,
				Quantity:    p.Quantity,
			})
		}
		orders = append(orders, &order.Order{
			OrderId: o.OrderId,
			Address: &order.Address{
				Name:          o.Address.Name,
				PhoneNumber:   o.Address.PhoneNumber,
				Province:      o.Address.Province,
				City:          o.Address.City,
				Region:        o.Address.Region,
				DetailAddress: o.Address.DetailAddress,
			},
			Products:  products,
			Cost:      o.Cost,
			Status:    o.Status,
			CreatedAt: o.CreatedAt,
		})
	}
	return &order.ListOrderResponse{
		StatusCode: listOrderResp.StatusCode,
		StatusMsg:  listOrderResp.StatusMsg,
		Orders:     orders,
	}, nil
}
