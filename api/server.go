package api

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/controllers"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/helpers"
	//"log"
	"os"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	//if err := godotenv.Load(); err != nil {
	//	log.Print("no .env file found")
	//}
}

func Run() {

	//var err error
	//godotenv.Load()

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	helpers.AutomaticEmail()

	// This is for testing, when done, do well to comment
	//seeder.Load(server.DB)

	apiPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Listening to port %s", apiPort)
	server.Run(apiPort)

}
