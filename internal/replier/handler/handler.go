package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello from server 2",
	})
}

func HandleQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "reply from server",
	})
}
