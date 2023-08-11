package models

import (
	"errors"
	"github.com/3AM-Developer/dae/database"
	"time"
)

var ErrOTUInvalid = errors.New("invalid or expired OTU")

type OneTimeURL struct {
	ID        uint `gorm:"primary_key"`
	Token     string
	Used      bool
	CreatedAt time.Time
}

func CreatePreSignedURL(otu *OneTimeURL) (err error) {
	err = database.DB.Create(otu).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOTUs(otu *[]OneTimeURL) (err error) {
	err = database.DB.Find(otu).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOTU(otu *OneTimeURL, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(otu).Error
	if err != nil {
		return err
	}
	return nil
}

func VerifyOTU(token string) error {
	var otu OneTimeURL
	if err := database.DB.Where("token = ? AND used = ?", token, false).First(&otu).Error; err != nil {
		return ErrOTUInvalid
	}

	otu.Used = true
	database.DB.Save(&otu)

	return nil
}
