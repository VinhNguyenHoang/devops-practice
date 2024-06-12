package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGETRequest(c *gin.Context) {
	log.Println("Received request")
	c.JSON(http.StatusOK, gin.H{
		"message": "reply from server Jerry",
	})
}
