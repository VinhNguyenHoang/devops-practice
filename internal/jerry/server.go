package jerry

import (
	"context"
	"cs/internal/jerry/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// TODO: move common service logic to lib
type Server struct {
	// address with port
	Address string

	router *gin.Engine
}

func (s *Server) Start() (err error) {
	tp, err := initTracer()
	if err != nil {
		return err
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	s.router = gin.Default()
	s.router.Use(otelgin.Middleware("jerry"))

	// TODO: put something name more meaningful
	s.router.GET("/api", handler.HandleGETRequest)

	err = s.router.Run(s.Address)
	if err != nil {
		return err
	}
	return nil
}

func initTracer() (*sdktrace.TracerProvider, error) {
	traceExporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint("jaeger-collector.istio-system.svc.cluster.local:4318"))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("jerry"),
	)

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp, nil
}
