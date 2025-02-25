package service

import (
	"context"
	product "douyin_mall/product/kitex_gen/product"
	"testing"
)

func TestDeleteCategory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteCategoryService(ctx)
	// init req and assert value

	req := &product.CategoryDeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
