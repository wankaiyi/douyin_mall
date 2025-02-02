package nacos

import (
	"douyin_mall/common/utils/env"
	"fmt"
	kitexRigistry "github.com/cloudwego/kitex/pkg/registry"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

func RegisterService() kitexRigistry.Registry {
	namingClient := GetNamingClient()
	r := registry.NewNacosRegistry(namingClient)
	return r
}

func GetNamingClient() naming_client.INamingClient {
	clientConfig, serverConfigs := GetNacosConfig()
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	return namingClient
}

func GetNacosConfig() (constant.ClientConfig, []constant.ServerConfig) {
	currentEnv, _ := env.GetString("env")
	var logLevel string
	if currentEnv == "dev" {
		logLevel = "debug"
	} else {
		logLevel = "info"
	}
	clientConfig := constant.ClientConfig{
		NamespaceId: "e45ccc29-3e7d-4275-917b-febc49052d58",
		TimeoutMs:   5000,
		LogLevel:    logLevel,
	}
	nacosIp, err := env.GetString("NACOS_ADDR")
	if err != nil {
		fmt.Println("nacos ip为空" + err.Error())
		panic(err)
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacosIp,
			Port:   8848,
		},
	}
	return clientConfig, serverConfigs
}
