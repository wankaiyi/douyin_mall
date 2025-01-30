package rpc

import (
	"douyin_mall/rpc/kitex_gen/auth/authservice"
	"douyin_mall/rpc/kitex_gen/payment/paymentservice"
	"douyin_mall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"time"

	//prometheus "github.com/kitex-contrib/monitor-prometheus"
	"sync"

	"douyin_mall/common/clientsuite"
	"douyin_mall/rpc/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	//prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var (
	AuthClient    authservice.Client
	UserClient    userservice.Client
	PaymentClient paymentservice.Client
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
			CurrentServiceName: "api",
		})
		initAuthClient()
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user-service", commonSuite, client.WithRPCTimeout(3*time.Second))
	if err != nil {
		klog.Fatal("init user client failed: ", err)
	}
}

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth-service", commonSuite, client.WithRPCTimeout(3*time.Second))
	if err != nil {
		klog.Fatal("init auth client failed: ", err)
	}
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product-service", commonSuite, client.WithRPCTimeout(3*time.Second))
	if err != nil {
		klog.Fatal("init auth client failed: ", err)
	}
}
