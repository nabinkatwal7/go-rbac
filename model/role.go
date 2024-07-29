package model

import (
	"github.com/nabinkatwal7/go-rbac/db"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

func CreateRole(Role *Role) (err error){
	err = db.Database.Create(Role).Error

	if err != nil {
		return err
	}
	return nil
}

func GetRoles(Role *[]Role) (err error){
	err = db.Database.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRole(Role *Role, id uint) (err error){
	err = db.Database.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRole(Role *Role) (err error){
	err = db.Database.Save(Role).Error
	if err != nil {
		return err
	}
	return nil
}