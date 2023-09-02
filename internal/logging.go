package internal

import (
	"context"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
	"os"
)

// RequestContextKey is the key used to store the request context in the context
const RequestContextKey = "requestContext"

// Logger is the global logger
var Logger = zerolog.New(os.Stdout).With().
	Caller().
	Timestamp().
	Logger().
	Hook(tracingHook{})

// tracingHook is a zerolog hook that adds traceId and spanId to the log output
type tracingHook struct{}

// Run adds traceId and spanId to the log output
func (h tracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	zctx := e.GetCtx()
	reqCtx := zctx.Value(RequestContextKey)
	if reqCtx == nil {
		return
	}
	span := trace.SpanFromContext(reqCtx.(context.Context))
	spanContext := span.SpanContext()
	e.Str("traceId", spanContext.TraceID().String()).
		Str("spanId", spanContext.SpanID().String())
}

// init initializes the logger
func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
