package models

import "time"

type Booking struct {
	ID           uint       `gorm:"primary_key;auto_increment" json:"id"`
	User         User       `json:"user"`
	UserID       uint       `json:"user_id"`
	Room         Room       `json:"room"`
	RoomID       uint       `json:"room_id"`
	TotalPerson  uint       `json:"total_person"`
	BookingTime  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"booking_time"`
	Noted        string     `gorm:"size:20" json:"noted"`
	CheckInTime  string     `gorm:"size:20" json:"check_in_time"`  // YYYY-MM-DD HH:MM:SS
	CheckOutTime string     `gorm:"size:20" json:"check_out_time"` // YYYY-MM-DD HH:MM:SS
	BeginTime    string     `gorm:"-" json:"-"`                    // HH:MM:SS only used for calculation
	EndTime      string     `gorm:"-" json:"-"`                    // HH:MM:SS only used for calculation
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"default:NULL" json:"-"`
}
