package seeder

import (
	"fmt"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/models"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Email:     "admin@gmail.com",
		Password:  "password123",
		Photo:     "",
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}, models.User{
		Email:     "user1@gmail.com",
		Password:  "password123",
		Photo:     "",
		Role:      "guest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	},
}

var rooms = []models.Room{
	models.Room{
		RoomName:     "Room Meeting 1",
		RoomCapacity: "20",
		Photo:        "",
		AvlTime:      nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	},
	models.Room{
		RoomName:     "Room Meeting 2",
		RoomCapacity: "12",
		Photo:        "",
		AvlTime:      nil,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	},
}

var bookings = []models.Booking{
	models.Booking{
		UserID:       2,
		User:         models.User{},
		RoomID:       1,
		Room:         models.Room{},
		TotalPerson:  10,
		BookingTime:  time.Now(),
		Noted:        "meeting pipeline project",
		CheckInTime:  time.Time{},
		CheckOutTime: time.Time{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	},
	models.Booking{
		UserID:       2,
		User:         models.User{},
		RoomID:       2,
		Room:         models.Room{},
		TotalPerson:  12,
		BookingTime:  time.Time{},
		Noted:        "meeting perubahan stack",
		CheckInTime:  time.Time{},
		CheckOutTime: time.Time{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	},
}

func Load(db *gorm.DB) {
	count := 0
	db.Debug().Model(&models.User{}).Count(&count)

	fmt.Println(count)

	if count == 0 {
		for i, _ := range users {
			err := db.Debug().Model(&models.User{}).Create(&users[i]).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)
			}

			err = db.Debug().Model(&models.Room{}).Create(&rooms[i]).Error
			if err != nil {
				log.Fatalf("cannot seed posts table: %v", err)
			}

			err = db.Debug().Model(&models.Booking{}).Create(&bookings[i]).Error
			if err != nil {
				log.Fatalf("cannot seed bookings table: %v", err)
			}
		}
	}

}
