package clientsuite

import (
	"douyin_mall/common/infra/nacos"
	sentinel "github.com/alibaba/sentinel-golang/pkg/adapters/kitex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

type CommonGrpcClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonGrpcClientSuite) Options() []client.Option {
	r := resolver.NewNacosResolver(nacos.GetNamingClient())
	opts := []client.Option{
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// sentinel客户端中间件
		client.WithMiddleware(sentinel.SentinelClientMiddleware()),
	}

	return opts
}
