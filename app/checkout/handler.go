package main

import (
	"context"
	"douyin_mall/checkout/biz/service"
	checkout "douyin_mall/checkout/kitex_gen/checkout"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}

// CheckoutProductItems implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) CheckoutProductItems(ctx context.Context, req *checkout.CheckoutProductItemsReq) (resp *checkout.CheckoutProductItemsResp, err error) {
	resp, err = service.NewCheckoutProductItemsService(ctx).Run(req)

	return resp, err
}
