package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/auth"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/helpers"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/mail"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateBooking(c *gin.Context) {
	errList = map[string]string{}

	fmt.Println("test")

	roomID := c.Param("id")
	rid, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	fmt.Println(uint(rid))

	// check the token
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	fmt.Println(uid)

	// check if the room is available:
	//room := models.Room{}
	//err = server.DB.Debug().Model(models.Room{}).Where("id = ?", rid).Take(&room).Error
	//if err != nil {
	//	errList["Unauthorized"] = "Unauthorized"
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"status": http.StatusUnauthorized,
	//		"error":  errList,
	//	})
	//	return
	//}

	body, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	booking := models.Booking{}
	err = json.Unmarshal(body, &booking)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	// enter the userid and the roomid. The comment body is automatically passed
	booking.UserID = uid
	booking.RoomID = uint(rid)

	booking.Prepare()
	booking.Validate(server.DB)
	isCapacityOkay, capacity := booking.RoomCapacity(server.DB)

	if isCapacityOkay == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "capacity max is " + strconv.Itoa(int(capacity)),
		})
		return
	}

	bookingCreated, err := booking.SaveBooking(server.DB)
	if err != nil {
		formattedError := helpers.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": bookingCreated,
	})

	mail.SendMail("chipxitro@gmail.com", "booking")
	//mail.SendMail2()
}
