package main

import (
	"context"
	"douyin_mall/order/biz/service"
	order "douyin_mall/order/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// GetOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	resp, err = service.NewGetOrderService(ctx).Run(req)

	return resp, err
}

// SmartSearchOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SmartSearchOrder(ctx context.Context, req *order.SmartSearchOrderReq) (resp *order.SmartSearchOrderResp, err error) {
	resp, err = service.NewSmartSearchOrderService(ctx).Run(req)

	return resp, err
}

// SmartPlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SmartPlaceOrder(ctx context.Context, req *order.SmartPlaceOrderReq) (resp *order.SmartPlaceOrderResp, err error) {
	resp, err = service.NewSmartPlaceOrderService(ctx).Run(req)

	return resp, err
}
