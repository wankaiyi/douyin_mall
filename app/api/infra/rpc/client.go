package rpc

import (
	"douyin_mall/common/middleware"
	"douyin_mall/rpc/kitex_gen/auth/authservice"
	"douyin_mall/rpc/kitex_gen/cart/cartservice"
	"douyin_mall/rpc/kitex_gen/checkout/checkoutservice"
	"douyin_mall/rpc/kitex_gen/order/orderservice"
	"douyin_mall/rpc/kitex_gen/payment/paymentservice"
	"douyin_mall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"time"

	"sync"

	"douyin_mall/common/clientsuite"
	"douyin_mall/rpc/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	AuthClient     authservice.Client
	UserClient     userservice.Client
	PaymentClient  paymentservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	OrderClient    orderservice.Client
	CheckoutClient checkoutservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = os.Getenv("NACOS_ADDR")
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: "api",
		})
		initAuthClient()
		initUserClient()
		initProductClient()
		initPaymentClient()
		initCartClient()
		InitOrderClient()
		InitCheckoutClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user-service", commonSuite, client.WithRPCTimeout(3*time.Second), client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init user client failed: ", err)
	}
}

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth-service", commonSuite, client.WithRPCTimeout(3*time.Second), client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init auth client failed: ", err)
	}
}
func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment-service", commonSuite, client.WithRPCTimeout(3*time.Second), client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init payment client failed: ", err)
	}
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product-service", commonSuite, client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init product client failed: ", err)
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart-service", commonSuite, client.WithRPCTimeout(3*time.Second), client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init cart client failed: ", err)
	}
}

func InitOrderClient() {
	OrderClient, err = orderservice.NewClient("order-service", commonSuite, client.WithRPCTimeout(5*time.Second), client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init order client failed: ", err)
	}
}

func InitCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout-service", commonSuite, client.WithRPCTimeout(5*time.Second), client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init checkout client failed: ", err)
	}
}
