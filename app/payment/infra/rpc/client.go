package rpc

import (
	"douyin_mall/common/middleware"
	"douyin_mall/user/conf"
	"douyin_mall/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"

	"sync"

	"douyin_mall/common/clientsuite"
	"douyin_mall/rpc/kitex_gen/cart/cartservice"
	"douyin_mall/rpc/kitex_gen/order/orderservice"
	"douyin_mall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
	UserClient    userservice.Client

	once         sync.Once
	err          error
	registryAddr string
	commonSuite  client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = os.Getenv("NACOS_ADDR")
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: conf.GetConf().Kitex.Service,
		})
		initProductClient()
		initOrderClient()
		initCartClient()
		initUserClient()

	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product-catalog-service", commonSuite, client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init product client failed: ", err)
	}
}
func initCartClient() {
	CartClient, err = cartservice.NewClient("cart-service", commonSuite, client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init cart client failed: ", err)
	}
}
func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order-service", commonSuite, client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init order client failed: ", err)
	}
}
func initUserClient() {
	UserClient, err = userservice.NewClient("user-service", commonSuite, client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init user client failed: ", err)
	}
}
