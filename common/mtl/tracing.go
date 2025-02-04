package mtl

import (
	"context"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"

	"github.com/cloudwego/kitex/server"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

var TracerProvider *tracesdk.TracerProvider

func InitTracing(serviceName string, endpoint string) {
	exporter, err := otlptracegrpc.New(context.Background())
	if err != nil {
		panic(err)
	}
	server.RegisterShutdownHook(func() {
		exporter.Shutdown(context.Background()) //nolint:errcheck
	})
	processor := tracesdk.NewBatchSpanProcessor(exporter)
	res, err := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)))
	if err != nil {
		res = resource.Default()
	}
	TracerProvider = tracesdk.NewTracerProvider(tracesdk.WithSpanProcessor(processor), tracesdk.WithResource(res))
	otel.SetTracerProvider(TracerProvider)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(endpoint),
		provider.WithEnableMetrics(false),
		provider.WithInsecure(),
	)
	server.RegisterShutdownHook(func() {
		p.Shutdown(context.Background())
	})
}
