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

		v1.POST("/register", server.CreateUser)
		v1.POST("/login", server.Login)

		v1.POST("/create-room", server.CreateRoom)
		//v1.POST("/cloud-storage-bucket", gcsbucket.HandleFileUploadToBucket)
	}
}
