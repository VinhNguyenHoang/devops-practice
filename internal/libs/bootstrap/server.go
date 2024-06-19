package bootstrap

import (
	"context"
	"log"

	"cs/internal/libs/metrics"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type HTTPServer struct {
	Name    string
	Address string
	router  *gin.Engine

	TraceCollectorEndpoint string
}

type GinHandler = func(*gin.Context)

func (s *HTTPServer) Start(endpointMap map[string]GinHandler) (err error) {
	tp, err := metrics.InitTracer(s.Name, s.TraceCollectorEndpoint)
	if err != nil {
		return err
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	s.router = gin.Default()
	s.router.Use(otelgin.Middleware(s.Name))

	s.registerAPIs(endpointMap)

	err = s.router.Run(s.Address)
	if err != nil {
		return err
	}
	return nil
}

func (s *HTTPServer) registerAPIs(endpoints map[string]GinHandler) {
	for endpoint, handler := range endpoints {
		s.router.GET(endpoint, handler)
	}
}
