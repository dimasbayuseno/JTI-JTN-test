package cmd

import (
	"context"
	"initial/configuration"
	"initial/configuration/configurationotel"
	"initial/delivery"
	serve "initial/delivery/server"
	"initial/infrastructure/logger"
)

func Execute() {
	// init zerolog global
	logger.Log(context.Background()).Info().Msg("Starting application")
	ctx := context.Background()
	if configuration.Env().Otel.EnableLogs {
		cleanupLogs, err := configurationotel.InitLogs()
		if err != nil {
			panic(err)
		}
		defer cleanupLogs(ctx)
		logger.NewHook()
	}

	if configuration.Env().Otel.EnableMetric {
		cleanupMetric, err := configurationotel.InitMetric()
		if err != nil {
			panic(err)
		}
		defer cleanupMetric(ctx)
	}

	if configuration.Env().Otel.EnableTracing {
		cleanupTracer, err := configurationotel.InitTracer()
		if err != nil {
			panic(err)
		}
		defer cleanupTracer(ctx)
	}

	container := delivery.SetupContainer()

	server := serve.ServeHttp(container)
	server.Listen(configuration.Viper().GetString("SERVER.PORT"))
}
