package main

import (
	"gorm.io/gorm"
	"jd_workout_golang/lib/database"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
}

func UpCreateUsersTable() error {
	err := database.Connection.Migrator().CreateTable(&User{})
	if err != nil {
		return err
	}

	return nil
}

func DownCreateUsersTable() error {
	err := database.Connection.Migrator().DropTable(&User{})
	if err != nil {
		return err
	}

	return nil
}
