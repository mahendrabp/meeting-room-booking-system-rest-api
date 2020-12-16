package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Booking struct {
	ID           uint       `gorm:"primary_key;auto_increment" json:"id"`
	User         User       `json:"user"`
	UserID       uint       `gorm:"not null" json:"user_id"`
	Room         Room       `json:"room"`
	RoomID       uint       `gorm:"not null" json:"room_id"`
	TotalPerson  uint       `gorm:"not null" json:"total_person"`
	BookingTime  time.Time  `gorm:"null" json:"booking_time"`
	Noted        string     `gorm:"size:255" json:"noted"`
	CheckInTime  time.Time  `gorm:"size:20;null;default:NULL" json:"check_in_time"`  // YYYY-MM-DD HH:MM:SS
	CheckOutTime time.Time  `gorm:"size:20;null;default:NULL" json:"check_out_time"` // YYYY-MM-DD HH:MM:SS
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"default:NULL" json:"-"`
}

func (b *Booking) Prepare() {
	b.ID = 0
	b.User = User{}
	b.Room = Room{}
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Booking) Validate(db *gorm.DB) {
	roomCapacity := db.Select("room_capacity").Find(&Room{})
	fmt.Println(roomCapacity)
	//// SELECT name, age FROM users;
}

func (b *Booking) RoomCapacity(db *gorm.DB) (bool, uint) {

	var room Room

	db.Model(&Room{}).Select("room_capacity").Where("id = ?", b.RoomID).First(&room)
	roomCapacity, _ := strconv.ParseUint(room.RoomCapacity, 10, 32)
	fmt.Println(roomCapacity)
	fmt.Println(b.TotalPerson)

	convert := uint(roomCapacity)
	if convert < b.TotalPerson {
		return false, convert
	}

	return true, convert

}

func (b *Booking) SaveBooking(db *gorm.DB) (*Booking, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return &Booking{}, err
	}

	if b.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", b.UserID).Take(&b.User).Error
		if err != nil {
			return &Booking{}, err
		}

		err = db.Debug().Model(&Room{}).Where("id = ?", b.RoomID).Take(&b.Room).Error
		if err != nil {
			return &Booking{}, err
		}
	}

	return b, nil
}
