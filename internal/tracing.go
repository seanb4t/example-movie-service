package internal

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var Tracer = otel.Tracer("github.com/seanb4t/example-movie-service")

func InitTracer() (*trace.TracerProvider, error) {
	r, err := resource.Merge(resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("Movie Service"),
		),
	)

	if err != nil {
		panic(err)
	}

	//exporter, err := stdout.New(stdout.WithPrettyPrint())
	//if err != nil {
	//	return nil, err
	//}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(r),
		//sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
