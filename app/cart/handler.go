package main

import (
	"context"
	"douyin_mall/cart/biz/service"
	cart "douyin_mall/cart/kitex_gen/cart"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	resp, err = service.NewEmptyCartService(ctx).Run(req)

	return resp, err
}

// SearchCarts implements the CartServiceImpl interface.
func (s *CartServiceImpl) SearchCarts(ctx context.Context, req *cart.SearchCartsReq) (resp *cart.SearchCartsResp, err error) {
	resp, err = service.NewSearchCartsService(ctx).Run(req)

	return resp, err
}

// InsertCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) InsertCart(ctx context.Context, req *cart.InsertCartReq) (resp *cart.InsertCartResp, err error) {
	resp, err = service.NewInsertCartService(ctx).Run(req)

	return resp, err
}

// SelectCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) SelectCart(ctx context.Context, req *cart.SelectCartReq) (resp *cart.SelectCartResp, err error) {
	resp, err = service.NewSelectCartService(ctx).Run(req)

	return resp, err
}

// DeleteCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) DeleteCart(ctx context.Context, req *cart.DeleteCartReq) (resp *cart.DeleteCartResp, err error) {
	resp, err = service.NewDeleteCartService(ctx).Run(req)

	return resp, err
}

// UpdateCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateCart(ctx context.Context, req *cart.UpdateCartReq) (resp *cart.UpdateCartResp, err error) {
	resp, err = service.NewUpdateCartService(ctx).Run(req)

	return resp, err
}
