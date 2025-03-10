// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	cart "douyin_mall/api/biz/router/cart"
	checkout "douyin_mall/api/biz/router/checkout"
	order "douyin_mall/api/biz/router/order"
	payment "douyin_mall/api/biz/router/payment"
	product "douyin_mall/api/biz/router/product"
	user "douyin_mall/api/biz/router/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	checkout.Register(r)

	order.Register(r)

	cart.Register(r)

	product.Register(r)

	payment.Register(r)

	user.Register(r)
}
