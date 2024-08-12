package model

import (
	"html"
	"strings"

	"github.com/nabinkatwal7/go-rbac/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	RoleID   uint   `gorm:"not null;default:3" json:"role_id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255; not null" json:"-"`
	Role     Role   `gorm:"constraing:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (user *User) Save() (*User, error) {
	err := db.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func GetUsers(User *[]User) (err error) {
	err = db.Database.Find(&User).Error

	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Database.Where("username = ?", username).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func GetUserByID(id uint) (User, error) {
	var user User

	err := db.Database.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUser(User *User, id int) (err error) {
	err = db.Database.Where("id = ?", id).First(&User).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(User *User) (err error) {
	err = db.Database.Omit("password").Updates(User).Error

	if err != nil {
		return err
	}
	return nil
}
