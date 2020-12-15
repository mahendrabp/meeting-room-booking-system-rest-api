package controllers

import (
	"fmt"
	//"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/helpers"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/models"
	//"io/ioutil"
	"net/http"
)

func (server *Server) CreateUser(c *gin.Context) {

	errList = map[string]string{}

	//body, err := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(body, err, "ini raw")
	//if err != nil {
	//	errList["Invalid_body"] = "Unable to get request"
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"status": http.StatusUnprocessableEntity,
	//		"error":  errList,
	//	})
	//	return
	//}

	user := models.User{}

	//err = json.Unmarshal(body, &user)
	//if err != nil {
	//	errList["Unmarshal_error"] = "Cannot unmarshal body"
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"status": http.StatusUnprocessableEntity,
	//		"error":  errList,
	//	})
	//	return
	//}

	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	//user.Photo = c.PostForm("photo")
	file, err := c.FormFile("photo")
	fmt.Println(file, err)

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
