package service

import (
	"context"
	"douyin_mall/product/biz/dal"
	"douyin_mall/product/conf"
	"douyin_mall/product/infra"
	product "douyin_mall/product/kitex_gen/product"
	"testing"
)

func TestUnlockProductQuantity_Run(t *testing.T) {
	conf.GetConf()
	ctx := context.Background()
	dal.Init()
	infra.Init()
	s := NewUnlockProductQuantityService(ctx)
	// init req and assert value

	req := &product.ProductUnLockQuantityRequest{
		Products: []*product.ProductUnLockQuantity{
			{
				ProductId: 104,
				Quantity:  1,
			},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
