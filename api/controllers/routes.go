package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) initializeRoutes() {
	v1 := server.Router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "connect")
		})

	}
}
