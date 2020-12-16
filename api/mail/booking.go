package mail

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendMail(emailUser, section string) {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}

	fmt.Println(os.Getenv("PASSWORD_EMAIL"))
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		emailUser,
	}

	var message []byte
	if section == "booking" {
		message = []byte("Subject: Notifikasi Booking!\r\n" + "\r\n" + "Anda baru saja membooking ruangan\r\n")
	} else if section == "check-in" {
		message = []byte("Subject: Notifikasi Checking!\r\n" + "\r\n" + "Anda sudah check-in\r\n")
	} else {
		message = []byte("Terima Kasih :)")
	}

	// smtp server configuration.
	smtpServer := smtpServer{host: os.Getenv("EMAIL_HOST"), port: os.Getenv("EMAIL_PORT")}

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err = smtp.SendMail(smtpServer.Address(), auth, "testbima121231@gmail.com", to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent!")
}
