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
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

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
	traceExporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	// traceExporter, err := stdout.New(stdout.WithPrettyPrint())
	// if err != nil {
	// 	return nil, err
	// }
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp, nil
}
