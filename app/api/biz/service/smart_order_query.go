package service

import (
	"context"
	order "douyin_mall/api/hertz_gen/api/order"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcOrder "douyin_mall/rpc/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type SmartOrderQueryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSmartOrderQueryService(Context context.Context, RequestContext *app.RequestContext) *SmartOrderQueryService {
	return &SmartOrderQueryService{RequestContext: RequestContext, Context: Context}
}

func (h *SmartOrderQueryService) Run(req *order.SmartOrderQueryRequest) (resp *order.SmartOrderQueryResponse, err error) {
	ctx := h.Context
	userId := ctx.Value(constant.UserId).(int32)
	smartSearchOrderResp, err := rpc.OrderClient.SmartSearchOrder(ctx, &rpcOrder.SmartSearchOrderReq{
		UserId:   userId,
		Uuid:     req.Uuid,
		Question: req.Question,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用AI查询订单失败，req: %v, error: %v", req, err)
		return nil, errors.New("服务器繁忙，请稍后再试")
	}

	var orderList []*order.Order
	for _, o := range smartSearchOrderResp.Orders {
		var productList []*order.Product
		for _, p := range o.Products {
			productList = append(productList, &order.Product{
				Id:          p.Id,
				Name:        p.Name,
				Description: p.Description,
				Picture:     p.Picture,
				Price:       p.Price,
				Quantity:    p.Quantity,
			})
		}
		orderList = append(orderList, &order.Order{
			OrderId: o.OrderId,
			Address: &order.Address{
				Name:          o.Address.Name,
				PhoneNumber:   o.Address.PhoneNumber,
				Province:      o.Address.Province,
				City:          o.Address.City,
				Region:        o.Address.Region,
				DetailAddress: o.Address.DetailAddress,
			},
			Products:  productList,
			Cost:      o.Cost,
			CreatedAt: o.CreatedAt,
			Status:    o.Status,
		})
	}

	return &order.SmartOrderQueryResponse{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Orders:     orderList,
	}, nil
}
