package rpc

import (
	"douyin_mall/rpc/kitex_gen/user/userservice"
	"douyin_mall/user/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"

	"sync"

	"douyin_mall/common/clientsuite"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient   userservice.Client
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
		initUserClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user-service", commonSuite)
	if err != nil {
		klog.Fatal("init user client failed: ", err)
	}
}
