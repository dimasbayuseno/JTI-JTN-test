package configurationotel

import (
	"context"
	"initial/configuration"
	"initial/infrastructure/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc/credentials"
)

func InitTracer() (func(context.Context) error, error) {
	ctx := context.Background()
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if configuration.Env().Otel.Insecure {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithEndpoint(configuration.Env().Otel.Endpoint),
			secureOption,
		),
	)

	if err != nil {
		logger.Log(ctx).Error().Err(err).Msg("could not initialize exporter")
		return nil, err
	}
	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName(configuration.Env().ServicName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		logger.Log(ctx).Error().Err(err).Msg("could not initialize resource")
		return nil, err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)

	if configuration.Env().Env == "development" {
		tracerProvider = sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithSyncer(exporter),
			sdktrace.WithResource(resources),
		)
	}

	otel.SetTracerProvider(
		tracerProvider,
	)
	return exporter.Shutdown, nil
}
