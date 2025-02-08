package main

import (
	"context"
	"douyin_mall/product/biz/service"
	product "douyin_mall/product/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// InsertProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) InsertProduct(ctx context.Context, req *product.InsertProductReq) (resp *product.InsertProductResp, err error) {
	resp, err = service.NewInsertProductService(ctx).Run(req)

	return resp, err
}

// SelectProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SelectProduct(ctx context.Context, req *product.SelectProductReq) (resp *product.SelectProductResp, err error) {
	resp, err = service.NewSelectProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// SelectCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SelectCategory(ctx context.Context, req *product.CategorySelectReq) (resp *product.CategorySelectResp, err error) {
	resp, err = service.NewSelectCategoryService(ctx).Run(req)

	return resp, err
}

// InsertCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) InsertCategory(ctx context.Context, req *product.CategoryInsertReq) (resp *product.CategoryInsertResp, err error) {
	resp, err = service.NewInsertCategoryService(ctx).Run(req)

	return resp, err
}

// DeleteCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteCategory(ctx context.Context, req *product.CategoryDeleteReq) (resp *product.CategoryDeleteResp, err error) {
	resp, err = service.NewDeleteCategoryService(ctx).Run(req)

	return resp, err
}

// UpdateCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateCategory(ctx context.Context, req *product.CategoryUpdateReq) (resp *product.CategoryUpdateResp, err error) {
	resp, err = service.NewUpdateCategoryService(ctx).Run(req)

	return resp, err
}
