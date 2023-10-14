package main

import (
	"jd_workout_golang/app/models"
	"jd_workout_golang/lib/database"
)


func UpAddImageToEuqipTable() error {
	err := database.Connection.Migrator().AddColumn(&models.Equip{}, "Image")
	if err != nil {
		return err
	}

	return nil
}

func DownAddImageToEuqipTable() error {
	err := database.Connection.Migrator().DropColumn(&models.Equip{}, "Image")
	if err != nil {
		return err
	}

	return nil
}