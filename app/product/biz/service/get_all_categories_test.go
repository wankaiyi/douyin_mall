package service

import (
	"context"
	product "douyin_mall/product/kitex_gen/product"
	"testing"
)

func TestGetAllCategories_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetAllCategoriesService(ctx)
	// init req and assert value

	req := &product.CategoryListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
