package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/gcsbucket"
	"net/http"
)

func (server *Server) initializeRoutes() {
	v1 := server.Router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "connect")
		})

		v1.POST("/register", server.CreateUser)
		v1.POST("/cloud-storage-bucket", gcsbucket.HandleFileUploadToBucket)
	}
}
