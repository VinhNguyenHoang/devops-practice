package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAsk(c *gin.Context) {
	resp, err := http.Get("http://localhost:30001/questions")
	if err != nil {
		log.Println(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(body),
	})
}
