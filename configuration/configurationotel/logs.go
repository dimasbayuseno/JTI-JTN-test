package configurationotel

import (
	"context"
	"initial/configuration"
	"initial/infrastructure/logger"

	otel "github.com/agoda-com/opentelemetry-logs-go"
	"github.com/agoda-com/opentelemetry-logs-go/exporters/otlp/otlplogs"
	"github.com/agoda-com/opentelemetry-logs-go/exporters/otlp/otlplogs/otlplogsgrpc"
	sdk "github.com/agoda-com/opentelemetry-logs-go/sdk/logs"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc/credentials"
)

func InitLogs() (func(context.Context) error, error) {
	ctx := context.Background()

	secureOption := otlplogsgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if configuration.Env().Otel.Insecure {
		secureOption = otlplogsgrpc.WithInsecure()
	}

	grpcClient := otlplogsgrpc.NewClient(
		otlplogsgrpc.WithEndpoint(configuration.Env().Otel.Endpoint),
		secureOption,
	)

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName(configuration.Env().ServicName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		logger.Log(ctx).Error().Err(err).Msg("logs: could not initialize resource")
		return nil, err
	}

	exporter, err := otlplogs.NewExporter(ctx,
		otlplogs.WithClient(grpcClient),
	)
	if err != nil {
		logger.Log(ctx).Error().Err(err).Msg("logs: could not initialize exporter")
		return nil, err
	}
	loggerProvider := sdk.NewLoggerProvider(
		sdk.WithBatcher(exporter),
		sdk.WithResource(resources),
	)

	otel.SetLoggerProvider(loggerProvider)

	return loggerProvider.Shutdown, nil
}
