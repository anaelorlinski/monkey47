package monkey47

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

func jaegerTracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(Component()),
			//attribute.String("environment", environment),
			//attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func InitJaegerOtel(url string) error {
	tracerProvider, err := jaegerTracerProvider(url)
	if err != nil {
		return err
	}

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}

// shortcut to otel tracer
func Tracer() trace.Tracer {
	return otel.Tracer(Component())
}

// TODO : add proper shutdown

// // Register our TracerProvider as the global so any imported
// // instrumentation in the future will default to using it.
// tracerProvider, err := monkey47.JaegerTracerProvider(JAEGER_URL)
// if err != nil {
// 	panic(err)
// }

// // TODO : hide in monkey47 module ?
// otel.SetTracerProvider(tracerProvider)
// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

// ctx, cancel := context.WithCancel(context.Background())
// defer cancel()

// // Cleanly shutdown and flush telemetry when the application exits.
// defer func(ctx context.Context) {
// 	// Do not make the application hang when it is shutdown.
// 	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
// 	defer cancel()
// 	if err := tracerProvider.Shutdown(ctx); err != nil {
// 		log.Fatal().Err(err)
// 	}
// }(ctx)
