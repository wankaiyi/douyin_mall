package service

import (
	"context"
	product "douyin_mall/product/kitex_gen/product"
	"testing"
)

func TestSelectBrand_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSelectBrandService(ctx)
	// init req and assert value

	req := &product.BrandSelectReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
