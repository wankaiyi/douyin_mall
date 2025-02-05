package service

import (
	"context"
	product "douyin_mall/product/kitex_gen/product"
	"testing"
)

func TestSelectProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSelectProductService(ctx)
	// init req and assert value

	req := &product.SelectProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
