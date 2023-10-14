package main

import (
	"jd_workout_golang/app/models"
	"jd_workout_golang/lib/database"
)

func UpAlterRestPasswordColumn() error {
	err := database.Connection.Migrator().AlterColumn(&models.User{}, "ResetPassword")
	if err != nil {
		return err
	}

	return nil
}

func DownAlterRestPasswordColumn() error {
	err := database.Connection.Migrator().AlterColumn(&models.User{}, "ResetPassword")
	if err != nil {
		return err
	}

	return nil
}