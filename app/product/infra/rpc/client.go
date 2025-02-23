package rpc

import (
	"douyin_mall/auth/conf"
	"douyin_mall/common/clientsuite"
	"github.com/cloudwego/kitex/client"
	"os"
	"sync"
)

var (
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

}
