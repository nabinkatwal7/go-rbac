package model

import (
	"github.com/nabinkatwal7/go-rbac/db"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	Name     string `gorm:"size:255;not null; unique" json:"name"`
	Location string `gorm:"size:255;not null" json:"location"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (room *Room) Save() (*Room, error) {
	err := db.Database.Create(&room).Error
	if err != nil {
		return &Room{}, err
	}
	return room, nil
}

func GetRooms(Room *[]Room) (err error) {
	err = db.Database.Find(&Room).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRoom(Room *Room, id int) (err error) {
	err = db.Database.Where("id = ?", id).First(&Room).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRoom(Room *Room) (err error) {
	err = db.Database.Save(Room).Error
	if err != nil {
		return err
	}
	return nil
}
