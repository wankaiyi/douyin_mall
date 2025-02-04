package serversuite

import (
	"douyin_mall/common/infra/nacos"
	"douyin_mall/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type CommonServerSuite struct {
	CurrentServiceName string
}

func (s CommonServerSuite) Options() []server.Option {
	r := nacos.RegisterService()
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithRegistry(r),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	}

	return opts
}
