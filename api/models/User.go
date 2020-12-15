package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Password  string     `gorm:"size:100;not null;" json:"password"`
	Photo     string     `gorm:"size:255;null;" json:"photo"`
	Role      string     `gorm:"size:10;not null" json:"role"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"default:NULL" json:"-"`
}
