package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

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

func (r *Room) Prepare() {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *Room) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if r.RoomName == "" {
		err = errors.New("Required Name")
		errorMessages["Required_Name"] = err.Error()
	}

	if r.RoomName != "" && len(r.RoomName) < 3 {
		err = errors.New("should be at least 3 characters")
		errorMessages["Invalid_Roomname"] = err.Error()
	}

	if r.RoomCapacity == "" {
		err = errors.New("Required Room Capacity")
		errorMessages["Required_Room_Capacity"] = err.Error()

	}
	return errorMessages
}

func (r *Room) CreateRoom(db *gorm.DB) (*Room, error) {
	var err error
	err = db.Debug().Model(&Room{}).Create(&r).Error
	if err != nil {
		return &Room{}, err
	}

	return r, nil
}

func (r *Room) FindAllRooms(db *gorm.DB) (*[]Room, error) {
	var err error
	var rooms []Room
	var bookings []Booking
	var roomId []uint

	//select room_id from bookings where booking_time today and check_out_time is null

	dt := time.Now()

	startDtFormatted := dt.Format("2006-01-02 00:00:00")
	endDtFormatted := dt.Format("2006-01-02") + " 23:59:59"

	err = db.Debug().Model(&Booking{}).
		Where("booking_time BETWEEN ? AND ?", startDtFormatted, endDtFormatted).
		Where("check_out_time is null").
		Find(&bookings).Error

	for _, b := range bookings {
		roomId = append(roomId, b.RoomID)
	}

	err = db.Debug().Not(roomId).Model(&Room{}).Limit(100).Order("created_at desc").Find(&rooms).Error

	//err = db.Model(&Room{}).Select("rooms.id, rooms.room_name, rooms.photo").Scan(&result{}).Error
	//
	//fmt.Println(result{})

	if err != nil {
		return &[]Room{}, err
	}

	return &rooms, nil
}

func (r *Room) FindRoomByID(db *gorm.DB, rid uint) (*Room, error) {
	var err error
	err = db.Debug().Model(&Room{}).Where("id = ?", rid).Take(&r).Error
	if err != nil {
		return &Room{}, err
	}

	return r, nil
}
