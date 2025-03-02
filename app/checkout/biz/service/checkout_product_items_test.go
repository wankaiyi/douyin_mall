package service

import (
	"context"
	checkout "douyin_mall/checkout/kitex_gen/checkout"
	"testing"
)

func TestCheckoutProductItems_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckoutProductItemsService(ctx)
	// init req and assert value

	req := &checkout.CheckoutProductItemsReq{
		UserId: 13,
		Items: []*checkout.ProductItem{
			{
				ProductId: 1,
				Quantity:  2,
			},
			{
				ProductId: 2,
				Quantity:  3,
			},
		},
		AddressId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
