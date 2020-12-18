package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Booking struct {
	ID           uint       `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint       `gorm:"not null;" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID" json:"user"`
	RoomID       uint       `gorm:";not null" json:"room_id"`
	Room         Room       `gorm:"foreignKey:RoomID" json:"room"`
	TotalPerson  uint       `gorm:"not null" json:"total_person"`
	BookingTime  time.Time  `gorm:"not null" json:"booking_time"`
	Noted        string     `gorm:"size:255" json:"noted"`
	CheckInTime  time.Time  `gorm:"default:NULL" json:"check_in_time"`  // YYYY-MM-DD HH:MM:SS
	CheckOutTime time.Time  `gorm:"default:NULL" json:"check_out_time"` // YYYY-MM-DD HH:MM:SS
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

func (b *Booking) GetAvailabilityRoom(db *gorm.DB, booking Booking, rid uint) bool {
	//var bookings []Booking

	startDtFormatted := booking.BookingTime.Format("2006-01-02 00:00:00")
	endDtFormatted := booking.BookingTime.Format("2006-01-02") + " 23:59:59"
	count := 0
	err := db.Debug().Model(&Booking{}).
		Where("booking_time BETWEEN ? AND ?", startDtFormatted, endDtFormatted).
		Where("room_id = ?", rid).
		Where("check_out_time is null").
		Count(&count).Error

	fmt.Println(count)
	fmt.Println(err)

	if err != nil {
		return false
	}

	if count != 0 {
		return false
	}
	return true
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

func (b *Booking) UpdateCheckIn(db *gorm.DB) (*Booking, error) {
	var err error
	fmt.Println(b)
	err = db.Debug().Model(&Booking{}).Where("id = ?", b.ID).Updates(Booking{CheckInTime: time.Now(), UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Booking{}, err
	}

	if b.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", b.UserID).Take(&b.User).Error
		if err != nil {
			return &Booking{}, err
		}
	}
	return b, nil
}

func (b *Booking) GetDetailBookTime(db *gorm.DB) ([]string, error) {

	var bookings []Booking
	var userId []uint
	var emailUser []User

	dt := time.Now()

	startDtFormatted := dt.Format("2006-01-02 00:00:00")
	endDtFormatted := dt.Format("2006-01-02") + " 23:59:59"

	err := db.Debug().Model(&Booking{}).
		Where("booking_time BETWEEN ? AND ?", startDtFormatted, endDtFormatted).
		Where("check_out_time is null").
		Find(&bookings).Error

	if err != nil {
		return []string{}, err
	}

	for _, u := range bookings {
		userId = append(userId, u.UserID)
	}

	db.Model(&User{}).Find(&emailUser, userId)

	var email []string
	for _, uid := range emailUser {
		email = append(email, uid.Email)
	}

	return email, nil
}
