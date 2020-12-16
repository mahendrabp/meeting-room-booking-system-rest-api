package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/middlewares"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/models"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var errList = make(map[string]string)

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	// If you are using mysql, i added support for you here(dont forgot to edit the .env file)
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("cant connect to %s database", Dbdriver)
			log.Fatal("error:", err)
		} else {
			fmt.Printf("connected to the %s database", Dbdriver)
		}
	} else {
		fmt.Println("Unknown Driver")
	}

	//database migration
	if isExist := server.DB.HasTable("bookings"); isExist != true {
		server.DB.Debug().AutoMigrate(&models.Booking{})
	}

	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Room{},
	)

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())

	server.initializeRoutes()

}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
