package spike

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExportEndpoints() map[string]func(c *gin.Context) {
	return map[string]func(c *gin.Context){
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
