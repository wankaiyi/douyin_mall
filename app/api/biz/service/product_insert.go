package service

import (
	"context"
	"douyin_mall/api/hertz_gen/api/product"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ProductInsertService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewProductInsertService(Context context.Context, RequestContext *app.RequestContext) *ProductInsertService {
	return &ProductInsertService{RequestContext: RequestContext, Context: Context}
}

func (h *ProductInsertService) Run(req *product.ProductInsertRequest) (resp *product.ProductInsertResponse, err error) {
	//1 先把参数解析出来
	productReq := rpcproduct.InsertProductReq{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryId:  req.CategoryId,
		BrandId:     req.BrandId,
	}
	//2 调用业务层的插入方法
	insertProduct, err := rpc.ProductClient.InsertProduct(h.Context, &productReq)
	if err != nil {
		return nil, err
	}
	//3 组装响应
	response := product.ProductInsertResponse{
		StatusCode: insertProduct.StatusCode,
		StatusMsg:  insertProduct.StatusMsg,
	}
	//4 返回响应
	return &response, nil
}
