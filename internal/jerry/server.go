package jerry

import (
	"cs/internal/jerry/handler"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type Server struct {
	// address with port
	Address string
}

func (s *Server) Start() error {
	err := router.Run(s.Address)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	router = gin.Default()

	// todo: put something name more meaningful
	router.GET("/api", handler.HandleGETRequest)
}
