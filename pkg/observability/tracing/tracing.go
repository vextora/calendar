package tracing

import (
	"context"
	"fmt"
	"oncomapi/pkg/config"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type Config struct {
	ServiceName string
	Environment string
	Protocol    string
	Endpoint    string
	Insecure    bool
	SampleRatio float64
}

func InitTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {

	tracerServiceName := config.GetEnvString(config.TracerServiceName)
	tracerEnv := config.GetEnvString(config.TracerEnv)
	tracerProtocol := config.GetEnvString(config.TracerProtocol)
	tracerEndpoint := config.GetEnvString(config.TracerEndpoint)

	tracerCfg := Config{
		ServiceName: tracerServiceName,
		Environment: tracerEnv,
		Protocol:    tracerProtocol,
		Endpoint:    tracerEndpoint,
		Insecure:    true,
		SampleRatio: 1.0,
	}

	var exporter sdktrace.SpanExporter
	var err error

	switch strings.ToLower(tracerCfg.Protocol) {
	case "http":
		exporter, err = otlptracehttp.New(ctx,
			otlptracehttp.WithEndpoint(tracerCfg.Endpoint),
			otlptracehttp.WithInsecure(),
		)
	case "grpc":
		exporter, err = otlptracegrpc.New(ctx,
			otlptracegrpc.WithEndpoint(tracerCfg.Endpoint),
			otlptracegrpc.WithInsecure(),
		)
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", tracerCfg.Protocol)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create exporter: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(tracerCfg.ServiceName),
			semconv.DeploymentEnvironment(tracerCfg.Environment),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(tracerCfg.SampleRatio)),
	)

	otel.SetTracerProvider(tp)

	return tp, nil
}
