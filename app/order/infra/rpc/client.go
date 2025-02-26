package rpc

import (
	"douyin_mall/order/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"

	"sync"

	"douyin_mall/common/clientsuite"
	"douyin_mall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
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
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product-service", commonSuite)
	if err != nil {
		klog.Fatal("init product client failed: ", err)
	}
}
