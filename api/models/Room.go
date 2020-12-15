package models

import "time"

type Room struct {
	ID           uint        `gorm:"primary_key;auto_increment" json:"id"`
	RoomName     string      `gorm:"size:100;not null;unique" json:"room_name"`
	RoomCapacity string      `gorm:"size:20;not null" json:"room_capacity"`
	Photo        string      `gorm:"size:255;null;" json:"photo"`
	AvlTime      []TimeSlice `gorm:"-" json:"avl_time"`
	CreatedAt    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time  `gorm:"default:NULL" json:"-"`
}
