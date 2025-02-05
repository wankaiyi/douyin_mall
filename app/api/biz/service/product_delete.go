package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	rpcproduct "douyin_mall/rpc/kitex_gen/product"

	product "douyin_mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ProductDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewProductDeleteService(Context context.Context, RequestContext *app.RequestContext) *ProductDeleteService {
	return &ProductDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *ProductDeleteService) Run(req *product.ProductDeleteRequest) (resp *product.ProductDeleteResponse, err error) {
	//1 先把参数解析出来
	productReq := rpcproduct.DeleteProductReq{
		Id: req.GetId(),
	}
	//2 调用业务层的删除方法
	deleteProduct, err := rpc.ProductClient.DeleteProduct(h.Context, &productReq)
	if err != nil {
		return nil, err
	}
	//3 组装响应
	return &product.ProductDeleteResponse{
		StatusCode: deleteProduct.StatusCode,
		StatusMsg:  deleteProduct.StatusMsg,
	}, nil
}
