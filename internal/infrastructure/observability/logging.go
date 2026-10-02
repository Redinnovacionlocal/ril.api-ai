package observability

import (
	"context"
	"log"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

func SetupLogging() *slog.Logger {
	var handler slog.Handler

	if os.Getenv("APP_ENV") == "local" {
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		base := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:       slog.LevelInfo,
			ReplaceAttr: gcpReplaceAttr,
		})
		handler = &gcpTraceHandler{
			Handler:   base,
			projectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
		}
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	log.SetFlags(0)
	log.SetOutput(slogWriter{logger})

	return logger
}

type gcpTraceHandler struct {
	slog.Handler
	projectID string
}

func (h *gcpTraceHandler) Handle(ctx context.Context, r slog.Record) error {
	if sc := trace.SpanContextFromContext(ctx); sc.IsValid() {
		if h.projectID != "" {
			r.AddAttrs(slog.String(
				"logging.googleapis.com/trace",
				"projects/"+h.projectID+"/traces/"+sc.TraceID().String(),
			))
		}
		r.AddAttrs(
			slog.String("logging.googleapis.com/spanId", sc.SpanID().String()),
			slog.Bool("logging.googleapis.com/trace_sampled", sc.IsSampled()),
		)
	}
	return h.Handler.Handle(ctx, r)
}

func (h *gcpTraceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &gcpTraceHandler{Handler: h.Handler.WithAttrs(attrs), projectID: h.projectID}
}

func (h *gcpTraceHandler) WithGroup(name string) slog.Handler {
	return &gcpTraceHandler{Handler: h.Handler.WithGroup(name), projectID: h.projectID}
}

func gcpReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		a.Key = "timestamp"
	case slog.MessageKey:
		a.Key = "message"
	case slog.LevelKey:
		a.Key = "severity"
		if lvl, ok := a.Value.Any().(slog.Level); ok {
			a.Value = slog.StringValue(gcpSeverity(lvl))
		}
	}
	return a
}

func gcpSeverity(l slog.Level) string {
	switch {
	case l >= slog.LevelError:
		return "ERROR"
	case l >= slog.LevelWarn:
		return "WARNING"
	case l >= slog.LevelInfo:
		return "INFO"
	default:
		return "DEBUG"
	}
}

type slogWriter struct{ l *slog.Logger }

func (w slogWriter) Write(p []byte) (int, error) {
	msg := string(p)
	if n := len(msg); n > 0 && msg[n-1] == '\n' {
		msg = msg[:n-1]
	}
	w.l.Info(msg)
	return len(p), nil
}
