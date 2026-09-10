package observability

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.36.0"
	adktelemetry "google.golang.org/adk/telemetry"
)

type ShutdownFunc func(context.Context) error

func SetupTracing(ctx context.Context) (ShutdownFunc, error) {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	if os.Getenv("OTEL_ENABLED") != "true" {
		slog.Info("tracing deshabilitado (OTEL_ENABLED != true)")
		return func(context.Context) error { return nil }, nil
	}

	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT")

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(envOr("APP_NAME", "ril-api")),
			semconv.ServiceVersion(envOr("APP_VERSION", "dev")),
			attribute.String("deployment.environment", envOr("APP_ENV", "unknown")),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creando resource de OpenTelemetry: %w", err)
	}

	opts := []adktelemetry.Option{adktelemetry.WithResource(res)}
	if otlpEndpoint != "" {
		slog.Info("tracing habilitado, exportando vía OTLP",
			slog.String("endpoint", otlpEndpoint),
			slog.String("service", envOr("APP_NAME", "ril-api")),
		)
	} else {
		opts = append(opts, adktelemetry.WithOtelToCloud(true))
		slog.Info("tracing habilitado, exportando a Google Cloud Trace",
			slog.String("service", envOr("APP_NAME", "ril-api")),
			slog.String("env", envOr("APP_ENV", "unknown")),
		)
	}

	providers, err := adktelemetry.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("inicializando telemetría del ADK: %w", err)
	}

	providers.SetGlobalOtelProviders()

	return providers.Shutdown, nil
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
