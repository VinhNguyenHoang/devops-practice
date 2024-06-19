package jerry

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cs/internal/libs/bootstrap"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

var tracer = otel.Tracer("test-tracer")

func ExportEndpoints() map[string]bootstrap.GinHandler {
	return map[string]bootstrap.GinHandler{
		"/hello": HandleRequest,
	}
}

func HandleRequest(c *gin.Context) {
	log.Println("Received request")

	reqCtx := c.Request.Context()

	// call to Spike
	spikeURL := os.Getenv("SPIKE_URL")
	if spikeURL == "" {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Missing Spike url"))
		return
	}

	c.Request.Header.Set("Content-Type", "application/json")
	req, err := http.NewRequest("GET", spikeURL+"/world", nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	otel.GetTextMapPropagator().Inject(reqCtx, propagation.HeaderCarrier(req.Header))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println("Header", req.Header)

	replyFromSpike := string(resBody)

	c.JSON(http.StatusOK, gin.H{
		"message": "reply from server Jerry:" + replyFromSpike,
	})
}
