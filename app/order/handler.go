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

// SearchOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SearchOrders(ctx context.Context, req *order.SearchOrdersReq) (resp *order.SearchOrdersResp, err error) {
	resp, err = service.NewSearchOrdersService(ctx).Run(req)

	return resp, err
}

// InsertOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) InsertOrder(ctx context.Context, req *order.InsertOrderReq) (resp *order.InsertOrderResp, err error) {
	resp, err = service.NewInsertOrderService(ctx).Run(req)

	return resp, err
}

// SelectOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SelectOrder(ctx context.Context, req *order.SelectOrderReq) (resp *order.SelectOrderResp, err error) {
	resp, err = service.NewSelectOrderService(ctx).Run(req)

	return resp, err
}

// DeleteOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, req *order.DeleteOrderReq) (resp *order.DeleteOrderResp, err error) {
	resp, err = service.NewDeleteOrderService(ctx).Run(req)

	return resp, err
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	resp, err = service.NewUpdateOrderService(ctx).Run(req)

	return resp, err
}
