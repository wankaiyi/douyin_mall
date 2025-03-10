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
	root.PUT("/category", append(_categoryinsertMw(), product.CategoryInsert)...)
	root.DELETE("/category", append(_categorydeleteMw(), product.CategoryDelete)...)
	root.GET("/category", append(_categoryselectMw(), product.CategorySelect)...)
	_category := root.Group("/category", _categoryMw()...)
	_category.POST("/update", append(_categoryupdateMw(), product.CategoryUpdate)...)
	root.DELETE("/product", append(_productdeleteMw(), product.ProductDelete)...)
	_product := root.Group("/product", _productMw()...)
	_product.POST("/list", append(_productselectlistMw(), product.ProductSelectList)...)
	_product.POST("/lockQuantity", append(_productlockquantityMw(), product.ProductLockQuantity)...)
	_product.POST("/update", append(_productupdateMw(), product.ProductUpdate)...)
	root.PUT("/product", append(_productinsertMw(), product.ProductInsert)...)
	_product0 := root.Group("/product", _product0Mw()...)
	_product0.GET("/:id", append(_productselectMw(), product.ProductSelect)...)
	{
		_product1 := root.Group("/product", _product1Mw()...)
		_product1.POST("/search", append(_searchMw(), product.Search)...)
	}
}
