package spike

import (
	"log"
	"net/http"

	"cs/internal/libs/bootstrap"

	"github.com/gin-gonic/gin"
)

func ExportEndpoints() map[string]bootstrap.GinHandler {
	return map[string]bootstrap.GinHandler{
		"/world": HandleRequest,
	}
}

func HandleRequest(c *gin.Context) {
	header := c.Request.Header
	log.Println("Header:", header)

	c.JSON(http.StatusOK, gin.H{
		"message": "reply from server Spike",
	})
}
