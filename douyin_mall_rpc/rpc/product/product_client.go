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
	SelectProductList(ctx context.Context, Req *product.SelectProductListReq, callOptions ...callopt.Option) (r *product.SelectProductListResp, err error)
	DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error)
	UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error)
	LockProductQuantity(ctx context.Context, Req *product.ProductLockQuantityRequest, callOptions ...callopt.Option) (r *product.ProductLockQuantityResponse, err error)
	SelectCategory(ctx context.Context, Req *product.CategorySelectReq, callOptions ...callopt.Option) (r *product.CategorySelectResp, err error)
	InsertCategory(ctx context.Context, Req *product.CategoryInsertReq, callOptions ...callopt.Option) (r *product.CategoryInsertResp, err error)
	DeleteCategory(ctx context.Context, Req *product.CategoryDeleteReq, callOptions ...callopt.Option) (r *product.CategoryDeleteResp, err error)
	UpdateCategory(ctx context.Context, Req *product.CategoryUpdateReq, callOptions ...callopt.Option) (r *product.CategoryUpdateResp, err error)
	GetAllCategories(ctx context.Context, Req *product.CategoryListReq, callOptions ...callopt.Option) (r *product.CategoryListResp, err error)
	SelectBrand(ctx context.Context, Req *product.BrandSelectReq, callOptions ...callopt.Option) (r *product.BrandSelectResp, err error)
	InsertBrand(ctx context.Context, Req *product.BrandInsertReq, callOptions ...callopt.Option) (r *product.BrandInsertResp, err error)
	DeleteBrand(ctx context.Context, Req *product.BrandDeleteReq, callOptions ...callopt.Option) (r *product.BrandDeleteResp, err error)
	UpdateBrand(ctx context.Context, Req *product.BrandUpdateReq, callOptions ...callopt.Option) (r *product.BrandUpdateResp, err error)
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

func (c *clientImpl) SelectProductList(ctx context.Context, Req *product.SelectProductListReq, callOptions ...callopt.Option) (r *product.SelectProductListResp, err error) {
	return c.kitexClient.SelectProductList(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) LockProductQuantity(ctx context.Context, Req *product.ProductLockQuantityRequest, callOptions ...callopt.Option) (r *product.ProductLockQuantityResponse, err error) {
	return c.kitexClient.LockProductQuantity(ctx, Req, callOptions...)
}

func (c *clientImpl) SelectCategory(ctx context.Context, Req *product.CategorySelectReq, callOptions ...callopt.Option) (r *product.CategorySelectResp, err error) {
	return c.kitexClient.SelectCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) InsertCategory(ctx context.Context, Req *product.CategoryInsertReq, callOptions ...callopt.Option) (r *product.CategoryInsertResp, err error) {
	return c.kitexClient.InsertCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteCategory(ctx context.Context, Req *product.CategoryDeleteReq, callOptions ...callopt.Option) (r *product.CategoryDeleteResp, err error) {
	return c.kitexClient.DeleteCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateCategory(ctx context.Context, Req *product.CategoryUpdateReq, callOptions ...callopt.Option) (r *product.CategoryUpdateResp, err error) {
	return c.kitexClient.UpdateCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) GetAllCategories(ctx context.Context, Req *product.CategoryListReq, callOptions ...callopt.Option) (r *product.CategoryListResp, err error) {
	return c.kitexClient.GetAllCategories(ctx, Req, callOptions...)
}

func (c *clientImpl) SelectBrand(ctx context.Context, Req *product.BrandSelectReq, callOptions ...callopt.Option) (r *product.BrandSelectResp, err error) {
	return c.kitexClient.SelectBrand(ctx, Req, callOptions...)
}

func (c *clientImpl) InsertBrand(ctx context.Context, Req *product.BrandInsertReq, callOptions ...callopt.Option) (r *product.BrandInsertResp, err error) {
	return c.kitexClient.InsertBrand(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteBrand(ctx context.Context, Req *product.BrandDeleteReq, callOptions ...callopt.Option) (r *product.BrandDeleteResp, err error) {
	return c.kitexClient.DeleteBrand(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateBrand(ctx context.Context, Req *product.BrandUpdateReq, callOptions ...callopt.Option) (r *product.BrandUpdateResp, err error) {
	return c.kitexClient.UpdateBrand(ctx, Req, callOptions...)
}
