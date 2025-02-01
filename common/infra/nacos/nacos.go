package nacos

import (
	kitexRigistry "github.com/cloudwego/kitex/pkg/registry"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"os"
)

func RegisterService() kitexRigistry.Registry {
	namingClient := GetNamingClient()
	r := registry.NewNacosRegistry(namingClient)
	return r
}

func GetNamingClient() naming_client.INamingClient {
	env := os.Getenv("env")
	var logLevel string
	if env == "dev" {
		logLevel = "debug"
	} else {
		logLevel = "info"
	}
	clientConfig := constant.ClientConfig{
		NamespaceId: "e45ccc29-3e7d-4275-917b-febc49052d58",
		TimeoutMs:   5000,
		LogLevel:    logLevel,
	}
	nacosIp := os.Getenv("NACOS_ADDR")
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacosIp,
			Port:   8848,
		},
	}
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	return namingClient
}
