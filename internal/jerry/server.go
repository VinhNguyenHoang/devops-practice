package jerry

import (
	"cs/internal/jerry/handler"

	"github.com/gin-gonic/gin"
)

func ExportEndpoints() map[string]func(c *gin.Context) {
	return map[string]func(c *gin.Context){
		"/api": handler.HandleGETRequest,
	}
}
