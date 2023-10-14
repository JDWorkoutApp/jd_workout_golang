package main

import (
	"jd_workout_golang/app/models"
	"jd_workout_golang/lib/database"
)

func UpAddResetPasswordColumnToUserTable() error {
	err := database.Connection.Migrator().AddColumn(&models.User{}, "ResetPassword")
	if err != nil {
		return err
	}

	return nil
}

func DownAddResetPasswordColumnToUserTable() error {
	err := database.Connection.Migrator().DropColumn(&models.User{}, "ResetPassword")
	if err != nil {
		return err
	}

	return nil
}