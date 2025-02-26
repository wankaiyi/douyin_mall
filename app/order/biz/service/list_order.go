package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/common/utils"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	"douyin_mall/order/infra/rpc"
	order "douyin_mall/order/kitex_gen/order"
	"douyin_mall/rpc/kitex_gen/product"
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

	orderIdList := make([]string, len(orderList))
	for i, o := range orderList {
		orderIdList[i] = o.OrderID
	}

	totalOrderItems, err := model.GetOrderItemsByOrderIdList(ctx, mysql.DB, orderIdList)
	if err != nil {
		klog.CtxErrorf(ctx, "数据库查询订单商品信息失败, error: %v", err)
		return nil, errors.WithStack(err)
	}

	orderItemsMap := make(map[string][]*model.OrderItem)
	productIdList := make([]int64, len(totalOrderItems))
	for _, item := range orderIdList {
		orderItemsMap[item] = make([]*model.OrderItem, 0)
	}
	for i, item := range totalOrderItems {
		productIdList[i] = int64(item.ProductID)
		if _, ok := orderItemsMap[item.OrderID]; ok {
			orderItemsMap[item.OrderID] = append(orderItemsMap[item.OrderID], item)
		}
	}

	productListReq := &product.SelectProductListReq{
		Ids: productIdList,
	}
	getProductListResp, err := rpc.ProductClient.SelectProductList(ctx, productListReq)
	if err != nil {
		klog.CtxErrorf(ctx, "rpc查询商品信息失败, req: %v, error: %v", productListReq, err)
		return nil, errors.WithStack(err)
	}
	productMap := make(map[int]*product.Product)
	for _, p := range getProductListResp.Products {
		productMap[int(p.Id)] = p
	}

	orders := make([]*order.Order, len(orderList))
	for i, o := range orderList {
		var products []*order.Product
		orderItems := orderItemsMap[o.OrderID]
		if orderItems == nil {
			continue
		}
		for _, item := range orderItems {
			p := productMap[int(item.ProductID)]
			if p == nil {
				continue
			}
			products = append(products, &order.Product{
				Id:       int32(p.Id),
				Name:     p.Name,
				Price:    float64(p.Price),
				Quantity: item.Quantity,
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
		}
	}
	return &order.ListOrderResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Orders:     orders,
	}, nil
}
