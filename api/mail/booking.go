package mail

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/models"
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

func messageEmail(bookingDetail *models.Booking, msg string) []byte {
	const AdminEmail = "bima@adminbooking.com"
	detailMsg := "From: " + AdminEmail + "\r\n" +
		"To: " + bookingDetail.User.Email + "\r\n"

	var message []byte
	if msg == "booking" {
		message = []byte(
			detailMsg +
				"Subject: Notifikasi Booking.\r\n" + "\r\n" +
				"Anda baru saja membooking ruangan " + bookingDetail.Room.RoomName +
				" untuk tanggal " + bookingDetail.BookingTime.Format("2006-01-02 15:04:05") + "\r\n")
	} else if msg == "check-in" {
		message = []byte(detailMsg + "Subject: Notifikasi Check-In.\r\n" +
			"\r\n" + "Anda sudah check-in pada " +
			bookingDetail.CheckInTime.Format("2006-01-02 15:04:05") + "\r\n")
	} else if msg == "reminder" {
		message = []byte(detailMsg + "Subject: Reminder Jadwal Booking.\r\n" + "\r\n" +
			"Jadwal Booking Anda Hari Ini di " + bookingDetail.Room.RoomName +
			" pada pukul " + bookingDetail.BookingTime.Format("2006-01-02 15:04:05") + "\r\n")
	} else {
		message = []byte(detailMsg + "Terima Kasih :)")
	}
	return message
}

//SendMail: using gmail as service sending email
func SendMail(bookingDetail *models.Booking, section string) {
	var err error

	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		bookingDetail.User.Email,
	}

	message := messageEmail(bookingDetail, section)
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

//func SendMailWithMailTrap(emailUser, section string) {
//
//	// Choose auth method and set it up
//	auth := smtp.PlainAuth("", os.Getenv("MAILTRAP_USER"), os.Getenv("MAILTRAP_PASSWORD"), "smtp.mailtrap.io")
//	fmt.Println(emailUser)
//	message := messageEmail(emailUser, section)
//
//	// Here we do it all: connect to our server, set up a message and send it
//	to := []string{emailUser}
//	msg := message
//	err := smtp.SendMail(os.Getenv("MAILTRAP_HOST"), auth, "bimadeveloper@mailtrap.io", to, msg)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	fmt.Println("Email with MailTrap Sent!")
//}
