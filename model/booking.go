package model

import (
	"github.com/nabinkatwal7/go-rbac/db"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"not null" json:"user_id"`
	RoomID uint   `gorm:"not null" json:"room_id"`
	Status string `gorm:"not null" json:"status"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Room   Room   `gorm:"constraing:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (booking *Booking) Save() (*Booking, error) {
	err := db.Database.Create(&booking).Error
	if err != nil {
		return &Booking{}, err
	}
	return booking, nil
}

func GetBookings(booking *[]Booking) (err error) {
	err = db.Database.Find(&booking).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserBookings(booking *[]Booking, userID uint) (err error) {
	err = db.Database.Where("user_id = ?", userID).Find(&booking).Error
	if err != nil {
		return err
	}
	return nil
}
