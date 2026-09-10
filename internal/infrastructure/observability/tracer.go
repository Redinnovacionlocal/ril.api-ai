package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const InstrumentationName = "ril.api-ia/app"

type Attrs map[string]any

type Span struct{ s trace.Span }

// Start abre un span hijo del que esté activo en ctx. Devolvé el ctx nuevo a lo
// que llames después y usá defer para cerrarlo:
//
//	ctx, span := observability.Start(ctx, "rag_proxy.generate_content", observability.Attrs{
//		"gen_ai.request.model": model,
//		"rag.query.length":     len(query),
//	})
//	defer span.End(&err) // registra el error si err != nil al salir (requiere named return)
func Start(ctx context.Context, name string, attrs Attrs) (context.Context, *Span) {
	ctx, s := otel.Tracer(InstrumentationName).Start(ctx, name)
	sp := &Span{s: s}
	sp.Set(attrs)
	return ctx, sp
}

// Set agrega atributos en cualquier momento (típicamente los resultados, una
// vez que terminó el trabajo).
func (sp *Span) Set(attrs Attrs) {
	if sp == nil || sp.s == nil || len(attrs) == 0 {
		return
	}
	sp.s.SetAttributes(toKV(attrs)...)
}

// Event deja una marca con hora dentro del span (ej: "empezó la descarga").
// Sirve para ver el reparto interno del tiempo sin abrir spans hijos.
func (sp *Span) Event(name string, attrs Attrs) {
	if sp == nil || sp.s == nil {
		return
	}
	sp.s.AddEvent(name, trace.WithAttributes(toKV(attrs)...))
}

// Fail marca el span como error (no lo cierra).
func (sp *Span) Fail(err error) {
	if sp == nil || sp.s == nil || err == nil {
		return
	}
	sp.s.RecordError(err)
	sp.s.SetStatus(codes.Error, err.Error())
}

// End cierra el span. Opcionalmente pasale el &err de la función (named return)
// y registra el error automáticamente si no es nil.
func (sp *Span) End(err ...*error) {
	if sp == nil || sp.s == nil {
		return
	}
	if len(err) > 0 && err[0] != nil {
		sp.Fail(*err[0])
	}
	sp.s.End()
}

// SetActive agrega atributos al span que ya esté activo en ctx (por ejemplo el
// execute_tool que crea el ADK). No abre uno nuevo. Si no hay span activo, no
// hace nada.
func SetActive(ctx context.Context, attrs Attrs) {
	(&Span{s: trace.SpanFromContext(ctx)}).Set(attrs)
}

func toKV(attrs Attrs) []attribute.KeyValue {
	kv := make([]attribute.KeyValue, 0, len(attrs))
	for k, v := range attrs {
		switch t := v.(type) {
		case string:
			kv = append(kv, attribute.String(k, t))
		case bool:
			kv = append(kv, attribute.Bool(k, t))
		case int:
			kv = append(kv, attribute.Int(k, t))
		case int64:
			kv = append(kv, attribute.Int64(k, t))
		case float64:
			kv = append(kv, attribute.Float64(k, t))
		case []string:
			kv = append(kv, attribute.StringSlice(k, t))
		}
	}
	return kv
}
