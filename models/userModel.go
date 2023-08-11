package models

import (
	"strings"

	"github.com/3AM-Developer/dae/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	RoleID   uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (user *User) Save() (*User, error) {
	err := database.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

// Gen encrypted Pass
func (user *User) BeforeSave(*gorm.DB) error {
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	user.Email = strings.TrimSpace(user.Email)

	return nil
}

// Get all users
func GetUsers(User *[]User) (err error) {
	err = database.DB.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by username
func GetUserByUsername(username string) (User, error) {
	var user User
	err := database.DB.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Validate user password
func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Get user by id
func GetUserById(id uint) (User, error) {
	var user User
	err := database.DB.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by id
func GetUser(User *User, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Update user
func UpdateUser(User *User) (err error) {
	err = database.DB.Omit("password").Updates(User).Error
	if err != nil {
		return err
	}
	return nil
}
