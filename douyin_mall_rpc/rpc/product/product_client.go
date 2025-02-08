package product

import (
	"context"
	product "douyin_mall/rpc/kitex_gen/product"

	"douyin_mall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
	InsertProduct(ctx context.Context, Req *product.InsertProductReq, callOptions ...callopt.Option) (r *product.InsertProductResp, err error)
	SelectProduct(ctx context.Context, Req *product.SelectProductReq, callOptions ...callopt.Option) (r *product.SelectProductResp, err error)
	DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error)
	UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error)
	SelectCategory(ctx context.Context, Req *product.CategorySelectReq, callOptions ...callopt.Option) (r *product.CategorySelectResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) InsertProduct(ctx context.Context, Req *product.InsertProductReq, callOptions ...callopt.Option) (r *product.InsertProductResp, err error) {
	return c.kitexClient.InsertProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SelectProduct(ctx context.Context, Req *product.SelectProductReq, callOptions ...callopt.Option) (r *product.SelectProductResp, err error) {
	return c.kitexClient.SelectProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SelectCategory(ctx context.Context, Req *product.CategorySelectReq, callOptions ...callopt.Option) (r *product.CategorySelectResp, err error) {
	return c.kitexClient.SelectCategory(ctx, Req, callOptions...)
}
