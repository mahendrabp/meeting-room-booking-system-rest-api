package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mahendrabp/meeting-room-booking-system-rest-api/api/helpers"
	"html"
	"strings"
	"time"
)

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

func (u *User) BeforeSave() error {
	hashedPassword, err := helpers.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Role = "guest"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
