package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/common/utils"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	order "douyin_mall/order/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	ctx := s.ctx
	userId := req.UserId
	orderList, err := model.GetOrdersByUserId(ctx, mysql.DB, userId)
	if err != nil {
		klog.CtxErrorf(ctx, "数据库查询订单信息失败, error: %v", err)
		return nil, errors.WithStack(err)
	}

	if orderList == nil {
		return &order.ListOrderResp{
			StatusCode: 0,
			StatusMsg:  constant.GetMsg(0),
		}, nil
	}

	orders := make([]*order.Order, len(orderList))
	for i, o := range orderList {
		var products []*order.Product
		orderItems := o.OrderItems
		for _, item := range orderItems {
			products = append(products, &order.Product{
				Id:          item.ProductID,
				Name:        item.ProductName,
				Price:       item.ProductPrice,
				Quantity:    item.Quantity,
				Picture:     item.ProductPicture,
				Description: item.ProductDescription,
			})
		}
		orders[i] = &order.Order{
			OrderId: o.OrderID,
			Address: &order.Address{
				Name:          o.Name,
				PhoneNumber:   o.PhoneNumber,
				Province:      o.Province,
				City:          o.City,
				Region:        o.Region,
				DetailAddress: o.DetailAddress,
			},
			Products:  products,
			Cost:      o.TotalCost,
			CreatedAt: utils.GetFormattedDateTime(o.CreatedAt),
			Status:    o.Status,
		}
	}
	return &order.ListOrderResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Orders:     orders,
	}, nil
}
