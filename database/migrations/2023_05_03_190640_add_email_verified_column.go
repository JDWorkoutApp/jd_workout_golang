package main

import (
	"jd_workout_golang/lib/database"
)

type user struct {
	EmailVerified int16 `json:"emailVerified" gorm:"default:0"`
}

func UpAddEmailVerifiedColumn() error {
	err := database.Connection.Migrator().AddColumn(&user{}, "EmailVerified")
	if err != nil {
		return err
	}

	return nil
}

func DownAddEmailVerifiedColumn() error {
	err := database.Connection.Migrator().DropColumn(&user{}, "EmailVerified")
	if err != nil {
		return err
	}

	return nil
}