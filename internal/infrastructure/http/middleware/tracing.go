package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.36.0"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "ril.api-ia/http"

func Tracing() gin.HandlerFunc {
	tracer := otel.Tracer(tracerName)
	propagator := otel.GetTextMapPropagator()

	return func(c *gin.Context) {
		ctx := propagator.Extract(
			c.Request.Context(),
			propagation.HeaderCarrier(c.Request.Header),
		)

		route := c.FullPath()
		if route == "" {
			route = "unknown"
		}
		spanName := c.Request.Method + " " + route

		ctx, span := tracer.Start(ctx, spanName,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				semconv.HTTPRequestMethodKey.String(c.Request.Method),
				semconv.HTTPRoute(route),
				semconv.URLPath(c.Request.URL.Path),
				semconv.UserAgentOriginal(c.Request.UserAgent()),
				attribute.String("http.client_ip", c.ClientIP()),
			),
		)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)

		start := time.Now()
		slog.InfoContext(ctx, "http request iniciada",
			slog.String("method", c.Request.Method),
			slog.String("route", route),
			slog.String("path", c.Request.URL.Path),
		)

		c.Next()

		status := c.Writer.Status()
		span.SetAttributes(semconv.HTTPResponseStatusCode(status))
		if status >= http.StatusInternalServerError {
			span.SetStatus(codes.Error, http.StatusText(status))
		}
		for _, e := range c.Errors {
			span.RecordError(e.Err)
		}

		slog.InfoContext(ctx, "http request finalizada",
			slog.String("method", c.Request.Method),
			slog.String("route", route),
			slog.Int("status", status),
			slog.Duration("duration", time.Since(start)),
		)
	}
}
