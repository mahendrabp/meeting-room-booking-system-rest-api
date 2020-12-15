package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/gcsbucket"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/helpers"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/models"
	"net/http"
)

func (server *Server) CreateRoom(c *gin.Context) {
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

	room := models.Room{}
	room.RoomName = c.PostForm("room_name")
	room.RoomCapacity = c.PostForm("room_capacity")
	room.Photo = path

	room.Prepare()
	errorMessages := room.Validate()
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	roomCreated, err := room.CreateRoom(server.DB)
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
		"response": roomCreated,
	})
}

func (server *Server) GetAvailableRoom(c *gin.Context) {
	room := models.Room{}

	rooms, err := room.FindAllRooms(server.DB)
	if err != nil {
		errList["No_room"] = "No Room Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": rooms,
	})
}
