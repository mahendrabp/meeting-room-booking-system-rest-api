package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/middlewares"
	"net/http"
)

func (server *Server) initializeRoutes() {
	v1 := server.Router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "connect")
		})

		v1.POST("/register", server.Register)
		v1.POST("/login", server.Login)

		v1.POST("/create-room", middlewares.TokenAuthMiddleware("admin"), server.CreateRoom)
		v1.GET("/available-room", middlewares.TokenAuthMiddleware(""), server.GetAvailableRoom)
		v1.GET("/available-room/:id", middlewares.TokenAuthMiddleware(""), server.GetRoom)

		v1.POST("/available-room/:id/booking", middlewares.TokenAuthMiddleware(""), server.CreateBooking)
		v1.PUT("/booking/:id", middlewares.TokenAuthMiddleware(""), server.UpdateCheckInTime)

		v1.GET("/secret-endpoint", server.AutomaticEmail)
	}
}
