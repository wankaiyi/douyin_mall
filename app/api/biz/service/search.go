package service

import (
	"context"
	"douyin_mall/api/biz/dal/mysql"
	"fmt"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchService(Context context.Context, RequestContext *app.RequestContext) *SearchService {
	return &SearchService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchService) Run(req *product.ProductRequest) (resp *product.ProductResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	db := mysql.DB
	var p []product.Product
	result := db.Table("tb_product").Select("*").Find(&p)
	products := make([]*product.Product, len(p))
	for i := range p {
		products[i] = &p[i]
	}

	fmt.Println(result)

	return &product.ProductResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Products:   products,
	}, nil
}
