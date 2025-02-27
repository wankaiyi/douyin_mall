package rpc

import (
	"douyin_mall/order/conf"
	"douyin_mall/rpc/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"

	"sync"

	"douyin_mall/common/clientsuite"
	"douyin_mall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
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
