package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/auth"
	"net/http"
)

func TokenAuthMiddleware(role string) gin.HandlerFunc {
	errList := make(map[string]string)
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			errList["unauthorized"] = "Unauthorized"
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errList,
			})
			c.Abort()
			return
		}

		if role == "admin" {
			payload, _ := auth.ExtractTokenRole(c.Request)
			if payload != "admin" {
				errList["unauthorized"] = "Unauthorized"
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  errList,
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
