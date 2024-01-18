package logger

import (
	"context"
	"initial/configuration"
	"initial/infrastructure/shared"
	"io"
	"os"
	"time"

	"github.com/agoda-com/opentelemetry-go/otelzerolog"
	otel "github.com/agoda-com/opentelemetry-logs-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Log(ctx context.Context) *zerolog.Logger {

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0770)

	}

	date := time.Now()
	logFile, _ := os.OpenFile("logs/log_"+date.Format("01-02-2006_15")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	logger := zerolog.New(multiWriter).With().Timestamp().Logger()
	if configuration.Viper().GetString("ENV") == "development" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	if val, ok := ctx.Value(shared.ContextKeyRequestID).(string); ok {
		logger = logger.With().Str(shared.ContextKeyRequestID.String(), val).Logger()
	}

	logger = logger.Hook(NewHook())
	log.Logger = logger

	return &logger
}

var otelZerologHook otelzerolog.Hook

func NewHook() otelzerolog.Hook {
	if otelZerologHook.Logger != nil {
		return otelZerologHook
	}

	logger := otel.GetLoggerProvider().Logger(
		configuration.Env().ServicName,
	)

	return otelzerolog.Hook{Logger: logger}
}
