package rpc

import (
	"douyin_mall/common/middleware"
	"douyin_mall/order/conf"
	"douyin_mall/rpc/kitex_gen/doubao_ai/doubaoaiservice"
	"douyin_mall/rpc/kitex_gen/payment/paymentservice"
	"douyin_mall/rpc/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"time"

	"sync"

	"douyin_mall/common/clientsuite"
	"douyin_mall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	DoubaoClient  doubaoaiservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	UserClient    userservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	commonSuite   client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = os.Getenv("NACOS_ADDR")
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: conf.GetConf().Kitex.Service,
		})
		initProductClient()
		InitPaymentClient()
		InitDoubaoClient()
		InitUserClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product-service", commonSuite)
	if err != nil {
		klog.Fatal("init product client failed: ", err)
	}
}

func InitPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment-service", commonSuite)
	if err != nil {
		klog.Fatal("init payment client failed: ", err)
	}
}

func InitDoubaoClient() {
	DoubaoClient, err = doubaoaiservice.NewClient("doubao-service", commonSuite, client.WithRPCTimeout(6*time.Second), client.WithMiddleware(middleware.ClientInterceptor))
	if err != nil {
		klog.Fatal("init doubao client failed: ", err)
	}
}

func InitUserClient() {
	UserClient, err = userservice.NewClient("user-service", commonSuite)
	if err != nil {
		klog.Fatal("init user client failed: ", err)
	}
}
