package configurationotel

import (
	"context"
	"initial/configuration"
	"initial/infrastructure/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc/credentials"
)

func InitMetric() (func(context.Context) error, error) {
	ctx := context.Background()
	secureOption := otlpmetricgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if configuration.Env().Otel.Insecure {
		secureOption = otlpmetricgrpc.WithInsecure()
	}

	exporter, err := otlpmetricgrpc.New(
		ctx,
		secureOption,
		otlpmetricgrpc.WithEndpoint(configuration.Env().Otel.Endpoint),
		otlpmetricgrpc.WithTemporalitySelector(preferDeltaTemporalitySelector),
	)

	if err != nil {
		logger.Log(ctx).Error().Err(err).Msg("could not initialize exporter")
		return nil, err
	}

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(configuration.Env().ServicName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		logger.Log(ctx).Error().Err(err).Msg("could not initialize resource")
		return nil, err
	}

	reader := sdkmetric.NewPeriodicReader(exporter)

	options := []sdkmetric.Option{
		sdkmetric.WithResource(resources),
		sdkmetric.WithReader(reader),
	}

	metricProvider := sdkmetric.NewMeterProvider(options...)
	otel.SetMeterProvider(metricProvider)

	return exporter.Shutdown, nil
}

func preferDeltaTemporalitySelector(kind sdkmetric.InstrumentKind) metricdata.Temporality {
	switch kind {
	case sdkmetric.InstrumentKindCounter,
		sdkmetric.InstrumentKindObservableCounter,
		sdkmetric.InstrumentKindHistogram:
		return metricdata.DeltaTemporality
	default:
		return metricdata.CumulativeTemporality
	}
}
