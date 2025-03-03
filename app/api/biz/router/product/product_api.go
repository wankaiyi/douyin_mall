// Code generated by hertz generator. DO NOT EDIT.

package product

import (
	product "douyin_mall/api/biz/handler/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/categories", append(_categoriesMw(), product.Categories)...)
	root.GET("/product", append(_productselectMw(), product.ProductSelect)...)
	root.DELETE("/product", append(_productdeleteMw(), product.ProductDelete)...)
	{
		_brand := root.Group("/brand", _brandMw()...)
		_brand.POST("/delete", append(_branddeleteMw(), product.BrandDelete)...)
		_brand.POST("/insert", append(_brandinsertMw(), product.BrandInsert)...)
		_brand.POST("/select", append(_brandselectMw(), product.BrandSelect)...)
		_brand.POST("/update", append(_brandupdateMw(), product.BrandUpdate)...)
	}
	{
		_category := root.Group("/category", _categoryMw()...)
		_category.POST("/delete", append(_categorydeleteMw(), product.CategoryDelete)...)
		_category.POST("/insert", append(_categoryinsertMw(), product.CategoryInsert)...)
		_category.POST("/select", append(_categoryselectMw(), product.CategorySelect)...)
		_category.POST("/update", append(_categoryupdateMw(), product.CategoryUpdate)...)
	}
	root.PUT("/product", append(_productinsertMw(), product.ProductInsert)...)
	_product := root.Group("/product", _productMw()...)
	_product.POST("/list", append(_productselectlistMw(), product.ProductSelectList)...)
	_product.POST("/lockQuantity", append(_productlockquantityMw(), product.ProductLockQuantity)...)
	_product.POST("/update", append(_productupdateMw(), product.ProductUpdate)...)
	{
		_product0 := root.Group("/product", _product0Mw()...)
		_product0.POST("/search", append(_searchMw(), product.Search)...)
	}
}
