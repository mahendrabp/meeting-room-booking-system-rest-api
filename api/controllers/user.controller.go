package controllers

import (
	"fmt"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/gcsbucket"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/helpers"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/models"
	//"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"net/http"
)

func (server *Server) CreateUser(c *gin.Context) {
	errList = map[string]string{}

	path, err := gcsbucket.HandleFileUploadToBucket(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	fmt.Println(path)

	user := models.User{}
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	user.Photo = path

	user.Prepare()
	errorMessages := user.Validate("register")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := helpers.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userCreated,
	})
}
