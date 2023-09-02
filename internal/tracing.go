package internal

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

var Tracer = otel.Tracer("github.com/seanb4t/example-movie-service")

func InitTracer() (*trace.TracerProvider, error) {
	//exporter, err := stdout.New(stdout.WithPrettyPrint())
	//if err != nil {
	//	return nil, err
	//}
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		//sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
